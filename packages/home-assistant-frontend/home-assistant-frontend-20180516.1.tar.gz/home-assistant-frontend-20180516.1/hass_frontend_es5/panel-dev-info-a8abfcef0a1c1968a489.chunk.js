(window.webpackJsonp=window.webpackJsonp||[]).push([[21],{165:function(e,n,t){"use strict";t(2),t(18),t(25),t(42);var o,r,a=t(4),i=t(0),s=(o=["\n    <style>\n      :host {\n        overflow: hidden; /* needed for text-overflow: ellipsis to work on ff */\n        @apply --layout-vertical;\n        @apply --layout-center-justified;\n        @apply --layout-flex;\n      }\n\n      :host([two-line]) {\n        min-height: var(--paper-item-body-two-line-min-height, 72px);\n      }\n\n      :host([three-line]) {\n        min-height: var(--paper-item-body-three-line-min-height, 88px);\n      }\n\n      :host > ::slotted(*) {\n        overflow: hidden;\n        text-overflow: ellipsis;\n        white-space: nowrap;\n      }\n\n      :host > ::slotted([secondary]) {\n        @apply --paper-font-body1;\n\n        color: var(--paper-item-body-secondary-color, var(--secondary-text-color));\n\n        @apply --paper-item-body-secondary;\n      }\n    </style>\n\n    <slot></slot>\n"],r=["\n    <style>\n      :host {\n        overflow: hidden; /* needed for text-overflow: ellipsis to work on ff */\n        @apply --layout-vertical;\n        @apply --layout-center-justified;\n        @apply --layout-flex;\n      }\n\n      :host([two-line]) {\n        min-height: var(--paper-item-body-two-line-min-height, 72px);\n      }\n\n      :host([three-line]) {\n        min-height: var(--paper-item-body-three-line-min-height, 88px);\n      }\n\n      :host > ::slotted(*) {\n        overflow: hidden;\n        text-overflow: ellipsis;\n        white-space: nowrap;\n      }\n\n      :host > ::slotted([secondary]) {\n        @apply --paper-font-body1;\n\n        color: var(--paper-item-body-secondary-color, var(--secondary-text-color));\n\n        @apply --paper-item-body-secondary;\n      }\n    </style>\n\n    <slot></slot>\n"],Object.freeze(Object.defineProperties(o,{raw:{value:Object.freeze(r)}})));
/**
@license
Copyright (c) 2015 The Polymer Project Authors. All rights reserved.
This code may only be used under the BSD style license found at http://polymer.github.io/LICENSE.txt
The complete set of authors may be found at http://polymer.github.io/AUTHORS.txt
The complete set of contributors may be found at http://polymer.github.io/CONTRIBUTORS.txt
Code distributed by Google as part of the polymer project is also
subject to an additional IP rights grant found at http://polymer.github.io/PATENTS.txt
*/
Object(a.a)({_template:Object(i.a)(s),is:"paper-item-body"})},171:function(e,n,t){"use strict";var o,r,a=t(0),i=t(3),s=(t(14),t(175),function(){function e(e,n){for(var t=0;t<n.length;t++){var o=n[t];o.enumerable=o.enumerable||!1,o.configurable=!0,"value"in o&&(o.writable=!0),Object.defineProperty(e,o.key,o)}}return function(n,t,o){return t&&e(n.prototype,t),o&&e(n,o),n}}()),l=(o=['\n    <ha-progress-button id="progress" progress="[[progress]]" on-click="buttonTapped"><slot></slot></ha-progress-button>\n'],r=['\n    <ha-progress-button id="progress" progress="[[progress]]" on-click="buttonTapped"><slot></slot></ha-progress-button>\n'],Object.freeze(Object.defineProperties(o,{raw:{value:Object.freeze(r)}})));var p=function(e){function n(){return function(e,n){if(!(e instanceof n))throw new TypeError("Cannot call a class as a function")}(this,n),function(e,n){if(!e)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return!n||"object"!=typeof n&&"function"!=typeof n?e:n}(this,(n.__proto__||Object.getPrototypeOf(n)).apply(this,arguments))}return function(e,n){if("function"!=typeof n&&null!==n)throw new TypeError("Super expression must either be null or a function, not "+typeof n);e.prototype=Object.create(n&&n.prototype,{constructor:{value:e,enumerable:!1,writable:!0,configurable:!0}}),n&&(Object.setPrototypeOf?Object.setPrototypeOf(e,n):e.__proto__=n)}(n,window.hassMixins.EventsMixin(i["a"])),s(n,[{key:"buttonTapped",value:function(){this.progress=!0;var e=this,n={domain:this.domain,service:this.service,serviceData:this.serviceData};this.hass.callService(this.domain,this.service,this.serviceData).then(function(){e.progress=!1,e.$.progress.actionSuccess(),n.success=!0},function(){e.progress=!1,e.$.progress.actionError(),n.success=!1}).then(function(){e.fire("hass-service-called",n)})}}],[{key:"template",get:function(){return Object(a.a)(l)}},{key:"properties",get:function(){return{hass:{type:Object},progress:{type:Boolean,value:!1},domain:{type:String},service:{type:String},serviceData:{type:Object,value:{}}}}}]),n}();customElements.define("ha-call-service-button",p)},172:function(e,n,t){"use strict";t(2),t(18),t(25),t(42),t(68);var o=document.createElement("template");o.setAttribute("style","display: none;"),o.innerHTML='<dom-module id="paper-dialog-shared-styles">\n  <template>\n    <style>\n      :host {\n        display: block;\n        margin: 24px 40px;\n\n        background: var(--paper-dialog-background-color, var(--primary-background-color));\n        color: var(--paper-dialog-color, var(--primary-text-color));\n\n        @apply --paper-font-body1;\n        @apply --shadow-elevation-16dp;\n        @apply --paper-dialog;\n      }\n\n      :host > ::slotted(*) {\n        margin-top: 20px;\n        padding: 0 24px;\n      }\n\n      :host > ::slotted(.no-padding) {\n        padding: 0;\n      }\n\n      \n      :host > ::slotted(*:first-child) {\n        margin-top: 24px;\n      }\n\n      :host > ::slotted(*:last-child) {\n        margin-bottom: 24px;\n      }\n\n      /* In 1.x, this selector was `:host > ::content h2`. In 2.x <slot> allows\n      to select direct children only, which increases the weight of this\n      selector, so we have to re-define first-child/last-child margins below. */\n      :host > ::slotted(h2) {\n        position: relative;\n        margin: 0;\n\n        @apply --paper-font-title;\n        @apply --paper-dialog-title;\n      }\n\n      /* Apply mixin again, in case it sets margin-top. */\n      :host > ::slotted(h2:first-child) {\n        margin-top: 24px;\n        @apply --paper-dialog-title;\n      }\n\n      /* Apply mixin again, in case it sets margin-bottom. */\n      :host > ::slotted(h2:last-child) {\n        margin-bottom: 24px;\n        @apply --paper-dialog-title;\n      }\n\n      :host > ::slotted(.paper-dialog-buttons),\n      :host > ::slotted(.buttons) {\n        position: relative;\n        padding: 8px 8px 8px 24px;\n        margin: 0;\n\n        color: var(--paper-dialog-button-color, var(--primary-color));\n\n        @apply --layout-horizontal;\n        @apply --layout-end-justified;\n      }\n    </style>\n  </template>\n</dom-module>',document.head.appendChild(o.content)},175:function(e,n,t){"use strict";t(59),t(99);var o,r,a=t(0),i=t(3),s=function(){function e(e,n){for(var t=0;t<n.length;t++){var o=n[t];o.enumerable=o.enumerable||!1,o.configurable=!0,"value"in o&&(o.writable=!0),Object.defineProperty(e,o.key,o)}}return function(n,t,o){return t&&e(n.prototype,t),o&&e(n,o),n}}(),l=(o=['\n    <style>\n      .container {\n        position: relative;\n        display: inline-block;\n      }\n\n      paper-button {\n        transition: all 1s;\n      }\n\n      .success paper-button {\n        color: white;\n        background-color: var(--google-green-500);\n        transition: none;\n      }\n\n      .error paper-button {\n        color: white;\n        background-color: var(--google-red-500);\n        transition: none;\n      }\n\n      paper-button[disabled] {\n        color: #c8c8c8;\n      }\n\n      .progress {\n        @apply --layout;\n        @apply --layout-center-center;\n        position: absolute;\n        top: 0;\n        left: 0;\n        right: 0;\n        bottom: 0;\n      }\n    </style>\n    <div class="container" id="container">\n      <paper-button id="button" disabled="[[computeDisabled(disabled, progress)]]" on-click="buttonTapped">\n        <slot></slot>\n      </paper-button>\n      <template is="dom-if" if="[[progress]]">\n        <div class="progress">\n          <paper-spinner active=""></paper-spinner>\n        </div>\n      </template>\n    </div>\n'],r=['\n    <style>\n      .container {\n        position: relative;\n        display: inline-block;\n      }\n\n      paper-button {\n        transition: all 1s;\n      }\n\n      .success paper-button {\n        color: white;\n        background-color: var(--google-green-500);\n        transition: none;\n      }\n\n      .error paper-button {\n        color: white;\n        background-color: var(--google-red-500);\n        transition: none;\n      }\n\n      paper-button[disabled] {\n        color: #c8c8c8;\n      }\n\n      .progress {\n        @apply --layout;\n        @apply --layout-center-center;\n        position: absolute;\n        top: 0;\n        left: 0;\n        right: 0;\n        bottom: 0;\n      }\n    </style>\n    <div class="container" id="container">\n      <paper-button id="button" disabled="[[computeDisabled(disabled, progress)]]" on-click="buttonTapped">\n        <slot></slot>\n      </paper-button>\n      <template is="dom-if" if="[[progress]]">\n        <div class="progress">\n          <paper-spinner active=""></paper-spinner>\n        </div>\n      </template>\n    </div>\n'],Object.freeze(Object.defineProperties(o,{raw:{value:Object.freeze(r)}})));var p=function(e){function n(){return function(e,n){if(!(e instanceof n))throw new TypeError("Cannot call a class as a function")}(this,n),function(e,n){if(!e)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return!n||"object"!=typeof n&&"function"!=typeof n?e:n}(this,(n.__proto__||Object.getPrototypeOf(n)).apply(this,arguments))}return function(e,n){if("function"!=typeof n&&null!==n)throw new TypeError("Super expression must either be null or a function, not "+typeof n);e.prototype=Object.create(n&&n.prototype,{constructor:{value:e,enumerable:!1,writable:!0,configurable:!0}}),n&&(Object.setPrototypeOf?Object.setPrototypeOf(e,n):e.__proto__=n)}(n,i["a"]),s(n,[{key:"tempClass",value:function(e){var n=this.$.container.classList;n.add(e),setTimeout(function(){n.remove(e)},1e3)}},{key:"ready",value:function(){var e=this;(function e(n,t,o){null===n&&(n=Function.prototype);var r=Object.getOwnPropertyDescriptor(n,t);if(void 0===r){var a=Object.getPrototypeOf(n);return null===a?void 0:e(a,t,o)}if("value"in r)return r.value;var i=r.get;return void 0!==i?i.call(o):void 0})(n.prototype.__proto__||Object.getPrototypeOf(n.prototype),"ready",this).call(this),this.addEventListener("click",function(n){return e.buttonTapped(n)})}},{key:"buttonTapped",value:function(e){this.progress&&e.stopPropagation()}},{key:"actionSuccess",value:function(){this.tempClass("success")}},{key:"actionError",value:function(){this.tempClass("error")}},{key:"computeDisabled",value:function(e,n){return e||n}}],[{key:"template",get:function(){return Object(a.a)(l)}},{key:"properties",get:function(){return{hass:{type:Object},progress:{type:Boolean,value:!1},disabled:{type:Boolean,value:!1}}}}]),n}();customElements.define("ha-progress-button",p)},178:function(e,n,t){"use strict";t(2),t(18);var o,r,a=t(74),i=(t(25),t(4)),s=t(0),l=(o=['\n    <style>\n\n      :host {\n        display: block;\n        @apply --layout-relative;\n      }\n\n      :host(.is-scrolled:not(:first-child))::before {\n        content: \'\';\n        position: absolute;\n        top: 0;\n        left: 0;\n        right: 0;\n        height: 1px;\n        background: var(--divider-color);\n      }\n\n      :host(.can-scroll:not(.scrolled-to-bottom):not(:last-child))::after {\n        content: \'\';\n        position: absolute;\n        bottom: 0;\n        left: 0;\n        right: 0;\n        height: 1px;\n        background: var(--divider-color);\n      }\n\n      .scrollable {\n        padding: 0 24px;\n\n        @apply --layout-scroll;\n        @apply --paper-dialog-scrollable;\n      }\n\n      .fit {\n        @apply --layout-fit;\n      }\n    </style>\n\n    <div id="scrollable" class="scrollable" on-scroll="updateScrollState">\n      <slot></slot>\n    </div>\n'],r=['\n    <style>\n\n      :host {\n        display: block;\n        @apply --layout-relative;\n      }\n\n      :host(.is-scrolled:not(:first-child))::before {\n        content: \'\';\n        position: absolute;\n        top: 0;\n        left: 0;\n        right: 0;\n        height: 1px;\n        background: var(--divider-color);\n      }\n\n      :host(.can-scroll:not(.scrolled-to-bottom):not(:last-child))::after {\n        content: \'\';\n        position: absolute;\n        bottom: 0;\n        left: 0;\n        right: 0;\n        height: 1px;\n        background: var(--divider-color);\n      }\n\n      .scrollable {\n        padding: 0 24px;\n\n        @apply --layout-scroll;\n        @apply --paper-dialog-scrollable;\n      }\n\n      .fit {\n        @apply --layout-fit;\n      }\n    </style>\n\n    <div id="scrollable" class="scrollable" on-scroll="updateScrollState">\n      <slot></slot>\n    </div>\n'],Object.freeze(Object.defineProperties(o,{raw:{value:Object.freeze(r)}})));
/**
@license
Copyright (c) 2015 The Polymer Project Authors. All rights reserved.
This code may only be used under the BSD style license found at http://polymer.github.io/LICENSE.txt
The complete set of authors may be found at http://polymer.github.io/AUTHORS.txt
The complete set of contributors may be found at http://polymer.github.io/CONTRIBUTORS.txt
Code distributed by Google as part of the polymer project is also
subject to an additional IP rights grant found at http://polymer.github.io/PATENTS.txt
*/
Object(i.a)({_template:Object(s.a)(l),is:"paper-dialog-scrollable",properties:{dialogElement:{type:Object}},get scrollTarget(){return this.$.scrollable},ready:function(){this._ensureTarget(),this.classList.add("no-padding")},attached:function(){this._ensureTarget(),requestAnimationFrame(this.updateScrollState.bind(this))},updateScrollState:function(){this.toggleClass("is-scrolled",this.scrollTarget.scrollTop>0),this.toggleClass("can-scroll",this.scrollTarget.offsetHeight<this.scrollTarget.scrollHeight),this.toggleClass("scrolled-to-bottom",this.scrollTarget.scrollTop+this.scrollTarget.offsetHeight>=this.scrollTarget.scrollHeight)},_ensureTarget:function(){this.dialogElement=this.dialogElement||this.parentElement,this.dialogElement&&this.dialogElement.behaviors&&this.dialogElement.behaviors.indexOf(a.b)>=0?(this.dialogElement.sizingTarget=this.scrollTarget,this.scrollTarget.classList.remove("fit")):this.dialogElement&&this.scrollTarget.classList.add("fit")}})},183:function(e,n,t){"use strict";t(2);var o,r,a=t(76),i=t(74),s=(t(172),t(4)),l=t(0),p=(o=['\n    <style include="paper-dialog-shared-styles"></style>\n    <slot></slot>\n'],r=['\n    <style include="paper-dialog-shared-styles"></style>\n    <slot></slot>\n'],Object.freeze(Object.defineProperties(o,{raw:{value:Object.freeze(r)}})));
/**
@license
Copyright (c) 2015 The Polymer Project Authors. All rights reserved.
This code may only be used under the BSD style license found at http://polymer.github.io/LICENSE.txt
The complete set of authors may be found at http://polymer.github.io/AUTHORS.txt
The complete set of contributors may be found at http://polymer.github.io/CONTRIBUTORS.txt
Code distributed by Google as part of the polymer project is also
subject to an additional IP rights grant found at http://polymer.github.io/PATENTS.txt
*/
Object(s.a)({_template:Object(l.a)(p),is:"paper-dialog",behaviors:[i.a,a.a],listeners:{"neon-animation-finish":"_onNeonAnimationFinish"},_renderOpened:function(){this.cancelAnimation(),this.playAnimation("entry")},_renderClosed:function(){this.cancelAnimation(),this.playAnimation("exit")},_onNeonAnimationFinish:function(){this.opened?this._finishRenderOpened():this._finishRenderClosed()}})},409:function(e,n,t){"use strict";t.r(n);t(119),t(117),t(82),t(131),t(178),t(183),t(51),t(165),t(73);var o,r,a=t(0),i=t(3),s=(t(171),t(118),t(132),t(62)),l=t(86),p=function(){function e(e,n){for(var t=0;t<n.length;t++){var o=n[t];o.enumerable=o.enumerable||!1,o.configurable=!0,"value"in o&&(o.writable=!0),Object.defineProperty(e,o.key,o)}}return function(n,t,o){return t&&e(n.prototype,t),o&&e(n,o),n}}(),c=function e(n,t,o){null===n&&(n=Function.prototype);var r=Object.getOwnPropertyDescriptor(n,t);if(void 0===r){var a=Object.getPrototypeOf(n);return null===a?void 0:e(a,t,o)}if("value"in r)return r.value;var i=r.get;return void 0!==i?i.call(o):void 0},d=(o=["\n    <style include=\"iron-positioning ha-style\">\n      :host {\n        -ms-user-select: initial;\n        -webkit-user-select: initial;\n        -moz-user-select: initial;\n      }\n\n      .content {\n        padding: 16px 0px 16px 0;\n      }\n\n      .about {\n        text-align: center;\n        line-height: 2em;\n      }\n\n      .version {\n        @apply --paper-font-headline;\n      }\n\n      .develop {\n        @apply --paper-font-subhead;\n      }\n\n      .about a {\n        color: var(--dark-primary-color);\n      }\n\n      .error-log-intro {\n        margin: 16px;\n      }\n\n      paper-icon-button {\n        float: right;\n      }\n\n      .error-log {\n        @apply --paper-font-code)\n        clear: both;\n        white-space: pre-wrap;\n        margin: 16px;\n      }\n\n      .system-log-intro {\n        margin: 16px;\n        border-top: 1px solid var(--light-primary-color);\n        padding-top: 16px;\n      }\n\n      paper-card {\n        display: block;\n        padding-top: 16px;\n      }\n\n      paper-item {\n        cursor: pointer;\n      }\n\n      .header {\n        @apply --paper-font-title;\n      }\n\n      paper-dialog {\n        border-radius: 2px;\n      }\n\n      @media all and (max-width: 450px), all and (max-height: 500px) {\n        paper-dialog {\n          margin: 0;\n          width: 100%;\n          max-height: calc(100% - 64px);\n\n          position: fixed !important;\n          bottom: 0px;\n          left: 0px;\n          right: 0px;\n          overflow: scroll;\n          border-bottom-left-radius: 0px;\n          border-bottom-right-radius: 0px;\n        }\n      }\n\n      .loading-container {\n        @apply --layout-vertical;\n        @apply --layout-center-center;\n        height: 100px;\n       }\n    </style>\n\n    <app-header-layout has-scrolling-region>\n      <app-header slot=\"header\" fixed>\n        <app-toolbar>\n          <ha-menu-button narrow='[[narrow]]' show-menu='[[showMenu]]'></ha-menu-button>\n          <div main-title>About</div>\n        </app-toolbar>\n      </app-header>\n\n      <div class='content'>\n        <div class='about'>\n          <p class='version'>\n            <a href='https://www.home-assistant.io'><img src=\"/static/icons/favicon-192x192.png\" height=\"192\" /></a><br />\n            Home Assistant<br />\n            [[hass.config.core.version]]\n          </p>\n          <p>\n            Path to configuration.yaml: [[hass.config.core.config_dir]]\n          </p>\n          <p class='develop'>\n            <a href='https://www.home-assistant.io/developers/credits/' target='_blank'>\n              Developed by a bunch of awesome people.\n            </a>\n          </p>\n          <p>\n            Published under the Apache 2.0 license<br />\n            Source:\n            <a href='https://github.com/home-assistant/home-assistant' target='_blank'>server</a> &mdash;\n            <a href='https://github.com/home-assistant/home-assistant-polymer' target='_blank'>frontend-ui</a>\n          </p>\n          <p>\n            Built using\n            <a href='https://www.python.org'>Python 3</a>,\n            <a href='https://www.polymer-project.org' target='_blank'>Polymer [[polymerVersion]]</a>,\n            Icons by <a href='https://www.google.com/design/icons/' target='_blank'>Google</a> and <a href='https://MaterialDesignIcons.com' target='_blank'>MaterialDesignIcons.com</a>.\n          </p>\n          <p>\n            Frontend JavaScript version: [[jsVersion]]\n            <template is='dom-if' if='[[customUiList.length]]'>\n              <div>\n                Custom UIs:\n                <template is='dom-repeat' items='[[customUiList]]'>\n                  <div>\n                    <a href='[[item.url]]' target='_blank'>[[item.name]]</a>: [[item.version]]\n                  </div>\n                </template>\n              </div>\n            </template>\n          </p>\n        </div>\n\n        <div class=\"system-log-intro\">\n          <paper-card>\n            <template is='dom-if' if='[[updating]]'>\n              <div class='loading-container'>\n                <paper-spinner active></paper-spinner>\n              </div>\n            </template>\n            <template is='dom-if' if='[[!updating]]'>\n              <template is='dom-if' if='[[!items.length]]'>\n                <div class='card-content'>There are no new issues!</div>\n              </template>\n              <template is='dom-repeat' items='[[items]]'>\n                <paper-item on-click='openLog'>\n                  <paper-item-body two-line>\n                    <div class=\"row\">\n                      [[item.message]]\n                    </div>\n                    <div secondary>\n                      [[formatTime(item.timestamp)]] [[item.source]] ([[item.level]])\n                    </div>\n                  </paper-item-body>\n                </paper-item>\n              </template>\n\n              <div class='card-actions'>\n                <ha-call-service-button\n                 hass='[[hass]]'\n                 domain='system_log'\n                 service='clear'\n                 >Clear</ha-call-service-button>\n                <ha-progress-button\n                 on-click='_fetchData'\n                 >Refresh</ha-progress-button>\n              </div>\n            </template>\n          </paper-card>\n        </div>\n        <p class='error-log-intro'>\n          Press the button to load the full Home Assistant log.\n          <paper-icon-button icon='mdi:refresh' on-click='refreshErrorLog'></paper-icon-button>\n        </p>\n        <div class='error-log'>[[errorLog]]</div>\n      </div>\n    </app-header-layout>\n\n    <paper-dialog with-backdrop id=\"showlog\">\n      <h2>Log Details ([[selectedItem.level]])</h2>\n      <paper-dialog-scrollable id=\"scrollable\">\n        <p>[[fullTimeStamp(selectedItem.timestamp)]]</p>\n        <template is='dom-if' if='[[selectedItem.message]]'>\n          <pre>[[selectedItem.message]]</pre>\n        </template>\n        <template is='dom-if' if='[[selectedItem.exception]]'>\n          <pre>[[selectedItem.exception]]</pre>\n        </template>\n      </paper-dialog-scrollable>\n    </paper-dialog>\n    "],r=["\n    <style include=\"iron-positioning ha-style\">\n      :host {\n        -ms-user-select: initial;\n        -webkit-user-select: initial;\n        -moz-user-select: initial;\n      }\n\n      .content {\n        padding: 16px 0px 16px 0;\n      }\n\n      .about {\n        text-align: center;\n        line-height: 2em;\n      }\n\n      .version {\n        @apply --paper-font-headline;\n      }\n\n      .develop {\n        @apply --paper-font-subhead;\n      }\n\n      .about a {\n        color: var(--dark-primary-color);\n      }\n\n      .error-log-intro {\n        margin: 16px;\n      }\n\n      paper-icon-button {\n        float: right;\n      }\n\n      .error-log {\n        @apply --paper-font-code)\n        clear: both;\n        white-space: pre-wrap;\n        margin: 16px;\n      }\n\n      .system-log-intro {\n        margin: 16px;\n        border-top: 1px solid var(--light-primary-color);\n        padding-top: 16px;\n      }\n\n      paper-card {\n        display: block;\n        padding-top: 16px;\n      }\n\n      paper-item {\n        cursor: pointer;\n      }\n\n      .header {\n        @apply --paper-font-title;\n      }\n\n      paper-dialog {\n        border-radius: 2px;\n      }\n\n      @media all and (max-width: 450px), all and (max-height: 500px) {\n        paper-dialog {\n          margin: 0;\n          width: 100%;\n          max-height: calc(100% - 64px);\n\n          position: fixed !important;\n          bottom: 0px;\n          left: 0px;\n          right: 0px;\n          overflow: scroll;\n          border-bottom-left-radius: 0px;\n          border-bottom-right-radius: 0px;\n        }\n      }\n\n      .loading-container {\n        @apply --layout-vertical;\n        @apply --layout-center-center;\n        height: 100px;\n       }\n    </style>\n\n    <app-header-layout has-scrolling-region>\n      <app-header slot=\"header\" fixed>\n        <app-toolbar>\n          <ha-menu-button narrow='[[narrow]]' show-menu='[[showMenu]]'></ha-menu-button>\n          <div main-title>About</div>\n        </app-toolbar>\n      </app-header>\n\n      <div class='content'>\n        <div class='about'>\n          <p class='version'>\n            <a href='https://www.home-assistant.io'><img src=\"/static/icons/favicon-192x192.png\" height=\"192\" /></a><br />\n            Home Assistant<br />\n            [[hass.config.core.version]]\n          </p>\n          <p>\n            Path to configuration.yaml: [[hass.config.core.config_dir]]\n          </p>\n          <p class='develop'>\n            <a href='https://www.home-assistant.io/developers/credits/' target='_blank'>\n              Developed by a bunch of awesome people.\n            </a>\n          </p>\n          <p>\n            Published under the Apache 2.0 license<br />\n            Source:\n            <a href='https://github.com/home-assistant/home-assistant' target='_blank'>server</a> &mdash;\n            <a href='https://github.com/home-assistant/home-assistant-polymer' target='_blank'>frontend-ui</a>\n          </p>\n          <p>\n            Built using\n            <a href='https://www.python.org'>Python 3</a>,\n            <a href='https://www.polymer-project.org' target='_blank'>Polymer [[polymerVersion]]</a>,\n            Icons by <a href='https://www.google.com/design/icons/' target='_blank'>Google</a> and <a href='https://MaterialDesignIcons.com' target='_blank'>MaterialDesignIcons.com</a>.\n          </p>\n          <p>\n            Frontend JavaScript version: [[jsVersion]]\n            <template is='dom-if' if='[[customUiList.length]]'>\n              <div>\n                Custom UIs:\n                <template is='dom-repeat' items='[[customUiList]]'>\n                  <div>\n                    <a href='[[item.url]]' target='_blank'>[[item.name]]</a>: [[item.version]]\n                  </div>\n                </template>\n              </div>\n            </template>\n          </p>\n        </div>\n\n        <div class=\"system-log-intro\">\n          <paper-card>\n            <template is='dom-if' if='[[updating]]'>\n              <div class='loading-container'>\n                <paper-spinner active></paper-spinner>\n              </div>\n            </template>\n            <template is='dom-if' if='[[!updating]]'>\n              <template is='dom-if' if='[[!items.length]]'>\n                <div class='card-content'>There are no new issues!</div>\n              </template>\n              <template is='dom-repeat' items='[[items]]'>\n                <paper-item on-click='openLog'>\n                  <paper-item-body two-line>\n                    <div class=\"row\">\n                      [[item.message]]\n                    </div>\n                    <div secondary>\n                      [[formatTime(item.timestamp)]] [[item.source]] ([[item.level]])\n                    </div>\n                  </paper-item-body>\n                </paper-item>\n              </template>\n\n              <div class='card-actions'>\n                <ha-call-service-button\n                 hass='[[hass]]'\n                 domain='system_log'\n                 service='clear'\n                 >Clear</ha-call-service-button>\n                <ha-progress-button\n                 on-click='_fetchData'\n                 >Refresh</ha-progress-button>\n              </div>\n            </template>\n          </paper-card>\n        </div>\n        <p class='error-log-intro'>\n          Press the button to load the full Home Assistant log.\n          <paper-icon-button icon='mdi:refresh' on-click='refreshErrorLog'></paper-icon-button>\n        </p>\n        <div class='error-log'>[[errorLog]]</div>\n      </div>\n    </app-header-layout>\n\n    <paper-dialog with-backdrop id=\"showlog\">\n      <h2>Log Details ([[selectedItem.level]])</h2>\n      <paper-dialog-scrollable id=\"scrollable\">\n        <p>[[fullTimeStamp(selectedItem.timestamp)]]</p>\n        <template is='dom-if' if='[[selectedItem.message]]'>\n          <pre>[[selectedItem.message]]</pre>\n        </template>\n        <template is='dom-if' if='[[selectedItem.exception]]'>\n          <pre>[[selectedItem.exception]]</pre>\n        </template>\n      </paper-dialog-scrollable>\n    </paper-dialog>\n    "],Object.freeze(Object.defineProperties(o,{raw:{value:Object.freeze(r)}})));var h=function(e){function n(){return function(e,n){if(!(e instanceof n))throw new TypeError("Cannot call a class as a function")}(this,n),function(e,n){if(!e)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return!n||"object"!=typeof n&&"function"!=typeof n?e:n}(this,(n.__proto__||Object.getPrototypeOf(n)).apply(this,arguments))}return function(e,n){if("function"!=typeof n&&null!==n)throw new TypeError("Super expression must either be null or a function, not "+typeof n);e.prototype=Object.create(n&&n.prototype,{constructor:{value:e,enumerable:!1,writable:!0,configurable:!0}}),n&&(Object.setPrototypeOf?Object.setPrototypeOf(e,n):e.__proto__=n)}(n,i["a"]),p(n,[{key:"ready",value:function(){var e=this;c(n.prototype.__proto__||Object.getPrototypeOf(n.prototype),"ready",this).call(this),this.addEventListener("hass-service-called",function(n){return e.serviceCalled(n)}),this.$.showlog.addEventListener("iron-overlay-opened",function(e){e.target.withBackdrop&&e.target.parentNode.insertBefore(e.target.backdropElement,e.target)})}},{key:"serviceCalled",value:function(e){e.detail.success&&"system_log"===e.detail.domain&&"clear"===e.detail.service&&(this.items=[])}},{key:"connectedCallback",value:function(){c(n.prototype.__proto__||Object.getPrototypeOf(n.prototype),"connectedCallback",this).call(this),this.$.scrollable.dialogElement=this.$.showlog,this._fetchData()}},{key:"refreshErrorLog",value:function(e){e&&e.preventDefault(),this.errorLog="Loading error log…",this.hass.callApi("GET","error_log").then(function(e){this.errorLog=e||"No errors have been reported."}.bind(this))}},{key:"fullTimeStamp",value:function(e){return new Date(1e3*e)}},{key:"formatTime",value:function(e){var n=(new Date).setHours(0,0,0,0),t=new Date(1e3*e);return new Date(1e3*e).setHours(0,0,0,0)<n?Object(s.a)(t):Object(l.a)(t)}},{key:"openLog",value:function(e){this.selectedItem=e.model.item,this.$.showlog.open()}},{key:"_fetchData",value:function(){this.updating=!0,this.hass.callApi("get","error/all").then(function(e){this.items=e,this.updating=!1}.bind(this))}}],[{key:"template",get:function(){return Object(a.a)(d)}},{key:"properties",get:function(){return{hass:Object,narrow:{type:Boolean,value:!1},showMenu:{type:Boolean,value:!1},polymerVersion:{type:String,value:Polymer.version},errorLog:{type:String,value:""},updating:{type:Boolean,value:!0},items:{type:Array,value:[]},selectedItem:Object,jsVersion:{type:String,value:"es5"},customUiList:{type:Array,value:window.CUSTOM_UI_LIST||[]}}}}]),n}();customElements.define("ha-panel-dev-info",h)}}]);