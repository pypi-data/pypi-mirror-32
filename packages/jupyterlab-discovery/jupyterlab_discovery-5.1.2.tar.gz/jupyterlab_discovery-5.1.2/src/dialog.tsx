// Copyright (c) Jupyter Development Team.
// Distributed under the terms of the Modified BSD License.

import {
  Dialog, showDialog
} from '@jupyterlab/apputils';

import * as React from 'react';


/**
 * Show a dialog box reporting an error during installation of an extension.
 *
 * @param name The name of the extension
 * @param errorMessage Any error message giving details about the failure.
 */
export
function reportInstallError(name: string, errorMessage?: string) {
  let entries = [];
  entries.push(<p>
    An error occurred installing <code>{name}</code>.
    </p>
  );
  if (errorMessage) {
    entries.push(<p>
      <span className="jp-discovery-dialog-subheader">
        Error message:
      </span>
      <pre>{errorMessage.trim()}</pre>
      </p>
    );
  }
  let body = (<div className="jp-discovery-dialog">
    {...entries}
    </div>
  );
  showDialog({
    title: 'Extension Installation Error',
    body,
    buttons: [
      Dialog.warnButton({label: 'OK'})
    ],
  });
}
