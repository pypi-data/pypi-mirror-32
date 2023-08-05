'use strict'

import {
  VDomModel
} from '@jupyterlab/apputils';

import {
  ServerConnection, ServiceManager, Kernel
} from '@jupyterlab/services';

import * as semver from 'semver';

import {
  doBuild
} from './build-helper';

import {
  Searcher, ISearchResult, IKernelInstallInfo
} from './query';

import {
  presentCompanions
} from './companions';

import {
  reportInstallError
} from './dialog';

/**
 * Information about an extension
 */
export
interface IEntry {
  name: string;
  description: string;
  installed: boolean;
  enabled: boolean;
  status: 'ok' | 'warning' | 'error' | 'deprecated' | null;
  latest_version: string;
  installed_version: string;
}


/**
 * Information about an installed extension
 */
export
interface IInstalledEntry {
  name: string;
  description: string;
  installed?: boolean;
  enabled: boolean;
  core: boolean;
  latest_version: string;
  installed_version: string;
  status: 'ok' | 'warning' | 'error' | 'deprecated' | null;
}

export
type KernelCompanion = {
  kernelInfo: IKernelInstallInfo,
  kernels: Kernel.ISpecModel[]
}


export
interface IActionReply {
  status: 'ok' | 'warning' | 'error' | null;
  message?: string;
}


/**
 * The server API path for querying/modifying installed extensions.
 */
const EXTENSION_API_PATH = "discovery/api/extensions"

/**
 * Extension actions that the server API accepts
 */
export
type Action = 'install' | 'uninstall' | 'enable' | 'disable';


/**
 * Model for an extension list.
 */
export
class ListModel extends VDomModel {
  constructor(serviceManager: ServiceManager) {
    super();
    this._installed = [];
    this._searchResult = [];
    this.serviceManager = serviceManager;
    this.serverConnectionSettings = ServerConnection.makeSettings();
  }

  /**
   * Utility function to check whether an entry can be updated.
   *
   * @param entry The entry to check.
   */
  static entryHasUpdate(entry: IEntry): boolean {
    if (!entry.installed || !entry.latest_version) {
      return false;
    }
    return semver.lt(entry.installed_version, entry.latest_version);
  }

  /**
   * A readonly array of the installed extensions.
   */
  get installed(): ReadonlyArray<IEntry> {
    return this._installed;
  }

  /**
   * A readonly array containing the latest search result
   */
  get searchResult(): ReadonlyArray<IEntry> {
    return this._searchResult;
  }

  /**
   * Translate search results from an npm repository query into entries.
   *
   * @param res Promise to an npm query result.
   */
  protected async translateSearchResult(res: Promise<ISearchResult>): Promise<{[key: string]: IEntry}> {
    let entries: {[key: string]: IEntry} = {};
    for (let obj of (await res).objects) {
      let pkg = obj.package;
      entries[pkg.name] = {
        name: pkg.name,
        description: pkg.description,
        installed: false,
        enabled: false,
        status: null,
        latest_version: pkg.version,
        installed_version: '',
      };
    }
    return entries;
  }

  /**
   * Translate installed extensions information from the server into entries.
   *
   * @param res Promise to the server reply data.
   */
  protected async translateInstalled(res: Promise<IInstalledEntry[]>): Promise<{[key: string]: IEntry}> {
    const promises = [];
    const entries: {[key: string]: IEntry} = {};
    for (let pkg of await res) {
      promises.push(res.then((info) => {
        entries[pkg.name] = {
          name: pkg.name,
          description: pkg.description,
          installed: pkg.installed !== false,
          enabled: pkg.enabled,
          status: pkg.status,
          latest_version: pkg.latest_version,
          installed_version: pkg.installed_version,
        };
      }));
    }
    return Promise.all(promises).then(() => {
      return entries;
    });
  }

  /**
   * The current NPM repository search query.
   *
   * Setting its value triggers a new search.
   */
  get query(): string {
    return this._query
  }
  set query(value: string) {
    this._query = value;
    this.update();
  }

  /**
   * The current NPM repository search page.
   *
   * The npm repository search is paginated by the `pagination` attribute.
   * The `page` value selects which page is used.
   *
   * Setting its value triggers a new search.
   */
  get page(): number {
    return this._page
  }
  set page(value: number) {
    this._page = value;
    this.update();
  }

  /**
   * The NPM repository search pagination.
   *
   * The npm repository search is paginated by the `pagination` attribute.
   * The `page` value selects which page is used.
   *
   * Setting its value triggers a new search.
   */
  get pagination(): number {
    return this._pagination
  }
  set pagination(value: number) {
    this._pagination = value;
    this.update();
  }

  /**
   * The total number of results in the current search.
   */
  get totalEntries(): number {
    return this._totalEntries;
  }

  /**
   * Make a request to the server for info about its installed extensions.
   */
  protected fetchInstalled(refreshInstalled=false): Promise<IInstalledEntry[]> {
    const url = new URL(EXTENSION_API_PATH, this.serverConnectionSettings.baseUrl);
    if (refreshInstalled) {
      url.searchParams.append('refresh', '1');
    }
    const request = ServerConnection.makeRequest(
      url.toString(), {}, this.serverConnectionSettings).then((response) => {
        handleError(response);
        return response.json() as Promise<IInstalledEntry[]>;
      });
    request.then(() => {
      this.serverConnectionError = null;
    }, (reason) => {
      this.serverConnectionError = reason.toString();
    });
    return request;
  }

  /**
   * Initialize the model.
   */
  initialize() {
    this.update().then(() => {
      this.initialized = true;
      this.stateChanged.emit(undefined);
    }).catch(() => {
      this.initialized = true;
      this.stateChanged.emit(undefined);
    });
  }

  /**
   * Update the current model.
   *
   * This will query the NPM repository, and the notebook server.
   *
   * Emits the `stateChanged` signal on succesfull completion.
   */
  protected async update(refreshInstalled=false) {
    let search = this.searcher.searchExtensions(this.query, this.page, this.pagination);
    let searchMapPromise = this.translateSearchResult(search);
    let installedMap = await this.translateInstalled(this.fetchInstalled(refreshInstalled));
    let searchMap;
    try {
      searchMap = await searchMapPromise;
      this.searchError = null;
    } catch (reason) {
      searchMap = {};
      this.searchError = reason.toString();
    }
    let installed: IEntry[] = [];
    for (let key of Object.keys(installedMap)) {
      installed.push(installedMap[key]);
    }
    this._installed = installed;

    let searchResult: IEntry[] = [];
    for (let key of Object.keys(searchMap)) {
      if (installedMap[key] === undefined) {
        searchResult.push(searchMap[key]);
      } else {
        searchResult.push(installedMap[key]);
      }
    }
    this._searchResult = searchResult;
    try {
      this._totalEntries = (await search).total;
    } catch (error) {
      this._totalEntries = 0;
    }
    this.stateChanged.emit(undefined);
  }

  /**
   * Send a request to the server to perform an action on an extension.
   *
   * @param action A valid action to perform.
   * @param entry The extension to perform the action on.
   */
  protected _performAction(action: string, entry: IEntry): Promise<IActionReply> {
    const url = new URL(EXTENSION_API_PATH, this.serverConnectionSettings.baseUrl);
    let request: RequestInit = {
      method: 'POST',
      body: JSON.stringify({
        cmd: action,
        extension_name: entry.name,
      }),
    };
    const completed = ServerConnection.makeRequest(url.toString(), request, this.serverConnectionSettings).then((response) => {
      handleError(response);
      this.triggerBuildCheck();
      return response.json() as Promise<IActionReply>;
    });
    completed.then(() => {
      this.serverConnectionError = null;
    }, (reason) => {
      this.serverConnectionError = reason.toString();
    });
    this._addPendingAction(completed);
    return completed;
  }

  protected _addPendingAction(pending: Promise<any>): void {
    // Add to pending actions collection
    this._pendingActions.push(pending);

    // Ensure action is removed when resolved
    const remove = () => {
      const i = this._pendingActions.indexOf(pending);
      this._pendingActions.splice(i, 1);
      this.stateChanged.emit(undefined);
    }
    pending.then(remove, remove);

    // Signal changed state
    this.stateChanged.emit(undefined);
  }

  hasPendingActions(): boolean {
    return this._pendingActions.length > 0;
  }

  /**
   * Install an extension.
   *
   * @param entry An entry indicating which extension to install.
   */
  install(entry: IEntry) {
    if (entry.installed) {
      // Updating
      return this._performAction('install', entry).then((data) => {
        if (data.status !== 'ok') {
          reportInstallError(entry.name, data.message);
        }
        this.update();
      });
    }
    this.checkCompanionPackages(entry).then((shouldInstall) => {
      if (shouldInstall) {
        return this._performAction('install', entry).then((data) => {
          if (data.status !== 'ok') {
            reportInstallError(entry.name, data.message);
          }
          this.update();
        });
      }
    })
  }

  /**
   * Uninstall an extension.
   *
   * @param entry An entry indicating which extension to uninstall.
   */
  uninstall(entry: IEntry) {
    if (!entry.installed) {
      throw new Error(`Not installed, cannot uninstall: ${entry.name}`);
    }
    this._performAction('uninstall', entry).then((data) => {
      this.update();
    });
  }

  /**
   * Enable an extension.
   *
   * @param entry An entry indicating which extension to enable.
   */
  enable(entry: IEntry) {
    if (entry.enabled) {
      throw new Error(`Already enabled: ${entry.name}`);
    }
    this._performAction('enable', entry).then((data) => {
      this.update();
    });
  }

  /**
   * Disable an extension.
   *
   * @param entry An entry indicating which extension to disable.
   */
  disable(entry: IEntry) {
    if (!entry.enabled) {
      throw new Error(`Already disabled: ${entry.name}`);
    }
    this._performAction('disable', entry).then((data) => {
      this.update();
    });
  }

  /**
   * Check for companion packages in kernels or server.
   *
   * @param entry An entry indicating which extension to check.
   */
  checkCompanionPackages(entry: IEntry): Promise<boolean> {
    return this.searcher.fetchPackageData(entry.name, entry.latest_version).then((data) => {
      if (!data || !data.jupyterlab || !data.jupyterlab.discovery) {
        return true;
      }
      let discovery = data.jupyterlab.discovery;
      let kernelCompanions: {kernelInfo: IKernelInstallInfo, kernels: Kernel.ISpecModel[]}[] = [];
      if (discovery.kernel) {
        // match specs
        for (let kernelInfo of discovery.kernel) {
          let matches = matchSpecs(kernelInfo, this.serviceManager.specs);
          kernelCompanions.push({kernelInfo, kernels: matches})
        }
      }
      if (kernelCompanions.length < 1 && !discovery.server) {
        return true;
      }
      return presentCompanions(kernelCompanions, discovery.server, this.serviceManager);
    });
  }

  /**
   * Trigger a build check to incorporate actions taken.
   */
  triggerBuildCheck(): void {
    let builder = this.serviceManager.builder;
    if (builder.isAvailable && !this.promptBuild) {
      const completed = builder.getStatus().then(response => {
        if (response.status === 'building') {
          // Piggy-back onto existing build
          // TODO: Can this cause dialog collision on build completion?
          return doBuild(builder);
        }
        if (response.status !== 'needed') {
          return;
        }
        if (!this.promptBuild) {
          this.promptBuild = true;
          this.stateChanged.emit(undefined);
        }
      });
      this._addPendingAction(completed);
    }
  }

  /**
   * Perform a build on the server
   */
  performBuild(): void {
    if (this.promptBuild) {
      this.promptBuild = false;
      this.stateChanged.emit(undefined);
    }
    const completed = doBuild(this.serviceManager.builder);
    this._addPendingAction(completed);
  }

  /**
   * Ignore a build recommendation
   */
  ignoreBuildRecommendation(): void {
    if (this.promptBuild) {
      this.promptBuild = false;
      this.stateChanged.emit(undefined);
    }
  }

  /**
   * Ignore a build recommendation
   */
  refreshInstalled(): void {
    const refresh = this.update(true);
    this._addPendingAction(refresh);
  }

  /**
   * Contains an error message if an error occurred when updating the model.
   */
  searchError: string | null = null;

  /**
   * Contains an error message if an error occurred when querying the server extension.
   */
  serverConnectionError: string | null = null;

  /**
   * Whether the model has finished async initialization.
   */
  initialized: boolean = false;

  /**
   * Whether a fresh build should be considered due to actions taken.
   */
  promptBuild: boolean = false;

  private _query: string = '';
  private _page: number = 0;
  private _pagination: number = 250;
  private _totalEntries: number = 0;

  protected _installed: IEntry[];
  protected _searchResult: IEntry[];
  protected _pendingActions: Promise<any>[] = [];

  /**
   * Settings for connecting to the notebook server.
   */
  protected serverConnectionSettings: ServerConnection.ISettings;

  /**
   * A helper for performing searches of jupyterlab extensions on the NPM repository.
   */
  protected searcher = new Searcher();

  protected serviceManager: ServiceManager;
}


/**
 * Match kernel specs against kernel spec regexps
 *
 * @param kernelInfo The info containing the regexp patterns
 * @param specs The available kernel specs.
 */
function matchSpecs(kernelInfo: IKernelInstallInfo, specs: Kernel.ISpecModels | null): Kernel.ISpecModel[] {
  if (!specs) {
    return [];
  }
  let matches: Kernel.ISpecModel[] = [];
  let reLang: RegExp | null = null;
  let reName: RegExp | null = null;
  if (kernelInfo.kernel_spec.language) {
    reLang = new RegExp(kernelInfo.kernel_spec.language);
  }
  if (kernelInfo.kernel_spec.display_name) {
    reName = new RegExp(kernelInfo.kernel_spec.display_name);
  }
  for (let key of Object.keys(specs.kernelspecs)) {
    let spec = specs.kernelspecs[key];
    let match = false;
    if (reLang) {
      match = reLang.test(spec.language);
    }
    if (!match && reName) {
      match = reName.test(spec.display_name);
    }
    if (match) {
      matches.push(spec);
      continue
    }
  }
  return matches;
}


function handleError(response: Response): Response {
  if (!response.ok) {
    throw new Error(`${response.status} (${response.statusText})`);
  }
  return response;
}
