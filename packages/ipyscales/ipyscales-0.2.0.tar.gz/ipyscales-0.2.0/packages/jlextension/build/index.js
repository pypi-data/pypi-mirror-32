"use strict";
// Copyright (c) Jupyter Development Team.
// Distributed under the terms of the Modified BSD License.
Object.defineProperty(exports, "__esModule", { value: true });
var coreutils_1 = require("@phosphor/coreutils");
var scales = require("jupyter-scales");
var base_1 = require("@jupyter-widgets/base");
var EXTENSION_ID = 'jupyter.extensions.jupyter-scales';
/**
 * The token identifying the JupyterLab plugin.
 */
exports.IJupyterScales = new coreutils_1.Token(EXTENSION_ID);
;
/**
 * The notebook diff provider.
 */
var scalesProvider = {
    id: EXTENSION_ID,
    requires: [base_1.IJupyterWidgetRegistry],
    activate: activateWidgetExtension,
    autoStart: true
};
exports.default = scalesProvider;
/**
 * Activate the widget extension.
 */
function activateWidgetExtension(app, widgetsManager) {
    widgetsManager.registerWidget({
        name: 'jupyter-scales',
        version: scales.JUPYTER_EXTENSION_VERSION,
        exports: scales,
    });
    return {};
}
//# sourceMappingURL=index.js.map