(window.webpackJsonp=window.webpackJsonp||[]).push([[10],{180:function(e,n,t){"use strict";t(3),t(20),t(23),t(46);var a,i,o=t(4),r=t(0),l=(a=["\n    <style>\n      :host {\n        overflow: hidden; /* needed for text-overflow: ellipsis to work on ff */\n        @apply --layout-vertical;\n        @apply --layout-center-justified;\n        @apply --layout-flex;\n      }\n\n      :host([two-line]) {\n        min-height: var(--paper-item-body-two-line-min-height, 72px);\n      }\n\n      :host([three-line]) {\n        min-height: var(--paper-item-body-three-line-min-height, 88px);\n      }\n\n      :host > ::slotted(*) {\n        overflow: hidden;\n        text-overflow: ellipsis;\n        white-space: nowrap;\n      }\n\n      :host > ::slotted([secondary]) {\n        @apply --paper-font-body1;\n\n        color: var(--paper-item-body-secondary-color, var(--secondary-text-color));\n\n        @apply --paper-item-body-secondary;\n      }\n    </style>\n\n    <slot></slot>\n"],i=["\n    <style>\n      :host {\n        overflow: hidden; /* needed for text-overflow: ellipsis to work on ff */\n        @apply --layout-vertical;\n        @apply --layout-center-justified;\n        @apply --layout-flex;\n      }\n\n      :host([two-line]) {\n        min-height: var(--paper-item-body-two-line-min-height, 72px);\n      }\n\n      :host([three-line]) {\n        min-height: var(--paper-item-body-three-line-min-height, 88px);\n      }\n\n      :host > ::slotted(*) {\n        overflow: hidden;\n        text-overflow: ellipsis;\n        white-space: nowrap;\n      }\n\n      :host > ::slotted([secondary]) {\n        @apply --paper-font-body1;\n\n        color: var(--paper-item-body-secondary-color, var(--secondary-text-color));\n\n        @apply --paper-item-body-secondary;\n      }\n    </style>\n\n    <slot></slot>\n"],Object.freeze(Object.defineProperties(a,{raw:{value:Object.freeze(i)}})));
/**
@license
Copyright (c) 2015 The Polymer Project Authors. All rights reserved.
This code may only be used under the BSD style license found at http://polymer.github.io/LICENSE.txt
The complete set of authors may be found at http://polymer.github.io/AUTHORS.txt
The complete set of contributors may be found at http://polymer.github.io/CONTRIBUTORS.txt
Code distributed by Google as part of the polymer project is also
subject to an additional IP rights grant found at http://polymer.github.io/PATENTS.txt
*/
Object(o.a)({_template:Object(r.a)(l),is:"paper-item-body"})},181:function(e,n,t){"use strict";t(3),t(182);var a,i,o=t(35),r=t(92),l=(t(119),t(118),t(117),t(4)),s=t(0),p=(a=['\n    <style>\n      :host {\n        display: block;\n      }\n\n      :host([hidden]) {\n        display: none !important;\n      }\n\n      label {\n        pointer-events: none;\n      }\n    </style>\n\n    <paper-input-container no-label-float$="[[noLabelFloat]]" always-float-label="[[_computeAlwaysFloatLabel(alwaysFloatLabel,placeholder)]]" auto-validate$="[[autoValidate]]" disabled$="[[disabled]]" invalid="[[invalid]]">\n\n      <label hidden$="[[!label]]" aria-hidden="true" for$="[[_inputId]]" slot="label">[[label]]</label>\n\n      <iron-autogrow-textarea class="paper-input-input" slot="input" id$="[[_inputId]]" aria-labelledby$="[[_ariaLabelledBy]]" aria-describedby$="[[_ariaDescribedBy]]" bind-value="{{value}}" invalid="{{invalid}}" validator$="[[validator]]" disabled$="[[disabled]]" autocomplete$="[[autocomplete]]" autofocus$="[[autofocus]]" inputmode$="[[inputmode]]" name$="[[name]]" placeholder$="[[placeholder]]" readonly$="[[readonly]]" required$="[[required]]" minlength$="[[minlength]]" maxlength$="[[maxlength]]" autocapitalize$="[[autocapitalize]]" rows$="[[rows]]" max-rows$="[[maxRows]]" on-change="_onChange"></iron-autogrow-textarea>\n\n      <template is="dom-if" if="[[errorMessage]]">\n        <paper-input-error aria-live="assertive" slot="add-on">[[errorMessage]]</paper-input-error>\n      </template>\n\n      <template is="dom-if" if="[[charCounter]]">\n        <paper-input-char-counter slot="add-on"></paper-input-char-counter>\n      </template>\n\n    </paper-input-container>\n'],i=['\n    <style>\n      :host {\n        display: block;\n      }\n\n      :host([hidden]) {\n        display: none !important;\n      }\n\n      label {\n        pointer-events: none;\n      }\n    </style>\n\n    <paper-input-container no-label-float\\$="[[noLabelFloat]]" always-float-label="[[_computeAlwaysFloatLabel(alwaysFloatLabel,placeholder)]]" auto-validate\\$="[[autoValidate]]" disabled\\$="[[disabled]]" invalid="[[invalid]]">\n\n      <label hidden\\$="[[!label]]" aria-hidden="true" for\\$="[[_inputId]]" slot="label">[[label]]</label>\n\n      <iron-autogrow-textarea class="paper-input-input" slot="input" id\\$="[[_inputId]]" aria-labelledby\\$="[[_ariaLabelledBy]]" aria-describedby\\$="[[_ariaDescribedBy]]" bind-value="{{value}}" invalid="{{invalid}}" validator\\$="[[validator]]" disabled\\$="[[disabled]]" autocomplete\\$="[[autocomplete]]" autofocus\\$="[[autofocus]]" inputmode\\$="[[inputmode]]" name\\$="[[name]]" placeholder\\$="[[placeholder]]" readonly\\$="[[readonly]]" required\\$="[[required]]" minlength\\$="[[minlength]]" maxlength\\$="[[maxlength]]" autocapitalize\\$="[[autocapitalize]]" rows\\$="[[rows]]" max-rows\\$="[[maxRows]]" on-change="_onChange"></iron-autogrow-textarea>\n\n      <template is="dom-if" if="[[errorMessage]]">\n        <paper-input-error aria-live="assertive" slot="add-on">[[errorMessage]]</paper-input-error>\n      </template>\n\n      <template is="dom-if" if="[[charCounter]]">\n        <paper-input-char-counter slot="add-on"></paper-input-char-counter>\n      </template>\n\n    </paper-input-container>\n'],Object.freeze(Object.defineProperties(a,{raw:{value:Object.freeze(i)}})));
/**
@license
Copyright (c) 2015 The Polymer Project Authors. All rights reserved.
This code may only be used under the BSD style license found at http://polymer.github.io/LICENSE.txt
The complete set of authors may be found at http://polymer.github.io/AUTHORS.txt
The complete set of contributors may be found at http://polymer.github.io/CONTRIBUTORS.txt
Code distributed by Google as part of the polymer project is also
subject to an additional IP rights grant found at http://polymer.github.io/PATENTS.txt
*/
Object(l.a)({_template:Object(s.a)(p),is:"paper-textarea",behaviors:[r.a,o.a],properties:{_ariaLabelledBy:{observer:"_ariaLabelledByChanged",type:String},_ariaDescribedBy:{observer:"_ariaDescribedByChanged",type:String},rows:{type:Number,value:1},maxRows:{type:Number,value:0}},get selectionStart(){return this.$.input.textarea.selectionStart},set selectionStart(e){this.$.input.textarea.selectionStart=e},get selectionEnd(){return this.$.input.textarea.selectionEnd},set selectionEnd(e){this.$.input.textarea.selectionEnd=e},_ariaLabelledByChanged:function(e){this._focusableElement.setAttribute("aria-labelledby",e)},_ariaDescribedByChanged:function(e){this._focusableElement.setAttribute("aria-describedby",e)},get _focusableElement(){return this.inputElement.textarea}})},182:function(e,n,t){"use strict";t(3);var a,i,o=t(11),r=(t(20),t(40)),l=t(4),s=t(0),p=t(2),d=(a=['\n    <style>\n      :host {\n        display: inline-block;\n        position: relative;\n        width: 400px;\n        border: 1px solid;\n        padding: 2px;\n        -moz-appearance: textarea;\n        -webkit-appearance: textarea;\n        overflow: hidden;\n      }\n\n      .mirror-text {\n        visibility: hidden;\n        word-wrap: break-word;\n        @apply --iron-autogrow-textarea;\n      }\n\n      .fit {\n        @apply --layout-fit;\n      }\n\n      textarea {\n        position: relative;\n        outline: none;\n        border: none;\n        resize: none;\n        background: inherit;\n        color: inherit;\n        /* see comments in template */\n        width: 100%;\n        height: 100%;\n        font-size: inherit;\n        font-family: inherit;\n        line-height: inherit;\n        text-align: inherit;\n        @apply --iron-autogrow-textarea;\n      }\n\n      textarea::-webkit-input-placeholder {\n        @apply --iron-autogrow-textarea-placeholder;\n      }\n\n      textarea:-moz-placeholder {\n        @apply --iron-autogrow-textarea-placeholder;\n      }\n\n      textarea::-moz-placeholder {\n        @apply --iron-autogrow-textarea-placeholder;\n      }\n\n      textarea:-ms-input-placeholder {\n        @apply --iron-autogrow-textarea-placeholder;\n      }\n    </style>\n\n    \x3c!-- the mirror sizes the input/textarea so it grows with typing --\x3e\n    \x3c!-- use &#160; instead &nbsp; of to allow this element to be used in XHTML --\x3e\n    <div id="mirror" class="mirror-text" aria-hidden="true">&nbsp;</div>\n\n    \x3c!-- size the input/textarea with a div, because the textarea has intrinsic size in ff --\x3e\n    <div class="textarea-container fit">\n      <textarea id="textarea" name$="[[name]]" aria-label$="[[label]]" autocomplete$="[[autocomplete]]" autofocus$="[[autofocus]]" inputmode$="[[inputmode]]" placeholder$="[[placeholder]]" readonly$="[[readonly]]" required$="[[required]]" disabled$="[[disabled]]" rows$="[[rows]]" minlength$="[[minlength]]" maxlength$="[[maxlength]]"></textarea>\n    </div>\n'],i=['\n    <style>\n      :host {\n        display: inline-block;\n        position: relative;\n        width: 400px;\n        border: 1px solid;\n        padding: 2px;\n        -moz-appearance: textarea;\n        -webkit-appearance: textarea;\n        overflow: hidden;\n      }\n\n      .mirror-text {\n        visibility: hidden;\n        word-wrap: break-word;\n        @apply --iron-autogrow-textarea;\n      }\n\n      .fit {\n        @apply --layout-fit;\n      }\n\n      textarea {\n        position: relative;\n        outline: none;\n        border: none;\n        resize: none;\n        background: inherit;\n        color: inherit;\n        /* see comments in template */\n        width: 100%;\n        height: 100%;\n        font-size: inherit;\n        font-family: inherit;\n        line-height: inherit;\n        text-align: inherit;\n        @apply --iron-autogrow-textarea;\n      }\n\n      textarea::-webkit-input-placeholder {\n        @apply --iron-autogrow-textarea-placeholder;\n      }\n\n      textarea:-moz-placeholder {\n        @apply --iron-autogrow-textarea-placeholder;\n      }\n\n      textarea::-moz-placeholder {\n        @apply --iron-autogrow-textarea-placeholder;\n      }\n\n      textarea:-ms-input-placeholder {\n        @apply --iron-autogrow-textarea-placeholder;\n      }\n    </style>\n\n    \x3c!-- the mirror sizes the input/textarea so it grows with typing --\x3e\n    \x3c!-- use &#160; instead &nbsp; of to allow this element to be used in XHTML --\x3e\n    <div id="mirror" class="mirror-text" aria-hidden="true">&nbsp;</div>\n\n    \x3c!-- size the input/textarea with a div, because the textarea has intrinsic size in ff --\x3e\n    <div class="textarea-container fit">\n      <textarea id="textarea" name\\$="[[name]]" aria-label\\$="[[label]]" autocomplete\\$="[[autocomplete]]" autofocus\\$="[[autofocus]]" inputmode\\$="[[inputmode]]" placeholder\\$="[[placeholder]]" readonly\\$="[[readonly]]" required\\$="[[required]]" disabled\\$="[[disabled]]" rows\\$="[[rows]]" minlength\\$="[[minlength]]" maxlength\\$="[[maxlength]]"></textarea>\n    </div>\n'],Object.freeze(Object.defineProperties(a,{raw:{value:Object.freeze(i)}})));
/**
@license
Copyright (c) 2015 The Polymer Project Authors. All rights reserved.
This code may only be used under the BSD style license found at http://polymer.github.io/LICENSE.txt
The complete set of authors may be found at http://polymer.github.io/AUTHORS.txt
The complete set of contributors may be found at http://polymer.github.io/CONTRIBUTORS.txt
Code distributed by Google as part of the polymer project is also
subject to an additional IP rights grant found at http://polymer.github.io/PATENTS.txt
*/
Object(l.a)({_template:Object(s.a)(d),is:"iron-autogrow-textarea",behaviors:[r.a,o.a],properties:{value:{observer:"_valueChanged",type:String,notify:!0},bindValue:{observer:"_bindValueChanged",type:String,notify:!0},rows:{type:Number,value:1,observer:"_updateCached"},maxRows:{type:Number,value:0,observer:"_updateCached"},autocomplete:{type:String,value:"off"},autofocus:{type:Boolean,value:!1},inputmode:{type:String},placeholder:{type:String},readonly:{type:String},required:{type:Boolean},minlength:{type:Number},maxlength:{type:Number},label:{type:String}},listeners:{input:"_onInput"},get textarea(){return this.$.textarea},get selectionStart(){return this.$.textarea.selectionStart},get selectionEnd(){return this.$.textarea.selectionEnd},set selectionStart(e){this.$.textarea.selectionStart=e},set selectionEnd(e){this.$.textarea.selectionEnd=e},attached:function(){navigator.userAgent.match(/iP(?:[oa]d|hone)/)&&(this.$.textarea.style.marginLeft="-3px")},validate:function(){var e=this.$.textarea.validity.valid;return e&&(this.required&&""===this.value?e=!1:this.hasValidator()&&(e=r.a.validate.call(this,this.value))),this.invalid=!e,this.fire("iron-input-validate"),e},_bindValueChanged:function(e){this.value=e},_valueChanged:function(e){var n=this.textarea;n&&(n.value!==e&&(n.value=e||0===e?e:""),this.bindValue=e,this.$.mirror.innerHTML=this._valueForMirror(),this.fire("bind-value-changed",{value:this.bindValue}))},_onInput:function(e){var n=Object(p.b)(e).path;this.value=n?n[0].value:e.target.value},_constrain:function(e){var n;for(e=e||[""],n=this.maxRows>0&&e.length>this.maxRows?e.slice(0,this.maxRows):e.slice(0);this.rows>0&&n.length<this.rows;)n.push("");return n.join("<br/>")+"&#160;"},_valueForMirror:function(){var e=this.textarea;if(e)return this.tokens=e&&e.value?e.value.replace(/&/gm,"&amp;").replace(/"/gm,"&quot;").replace(/'/gm,"&#39;").replace(/</gm,"&lt;").replace(/>/gm,"&gt;").split("\n"):[""],this._constrain(this.tokens)},_updateCached:function(){this.$.mirror.innerHTML=this._constrain(this.tokens)}})},189:function(e,n,t){"use strict";t(3);var a,i,o=t(93),r=t(74),l=(t(137),t(4)),s=t(0),p=(a=['\n    <style include="paper-dialog-shared-styles"></style>\n    <slot></slot>\n'],i=['\n    <style include="paper-dialog-shared-styles"></style>\n    <slot></slot>\n'],Object.freeze(Object.defineProperties(a,{raw:{value:Object.freeze(i)}})));
/**
@license
Copyright (c) 2015 The Polymer Project Authors. All rights reserved.
This code may only be used under the BSD style license found at http://polymer.github.io/LICENSE.txt
The complete set of authors may be found at http://polymer.github.io/AUTHORS.txt
The complete set of contributors may be found at http://polymer.github.io/CONTRIBUTORS.txt
Code distributed by Google as part of the polymer project is also
subject to an additional IP rights grant found at http://polymer.github.io/PATENTS.txt
*/
Object(l.a)({_template:Object(s.a)(p),is:"paper-dialog",behaviors:[r.a,o.a],listeners:{"neon-animation-finish":"_onNeonAnimationFinish"},_renderOpened:function(){this.cancelAnimation(),this.playAnimation("entry")},_renderClosed:function(){this.cancelAnimation(),this.playAnimation("exit")},_onNeonAnimationFinish:function(){this.opened?this._finishRenderOpened():this._finishRenderClosed()}})},398:function(e,n,t){"use strict";t.r(n);t(135),t(133),t(81),t(50),t(153),t(189),t(181),t(180),t(56);var a,i,o=t(0),r=t(1),l=(t(134),t(115),t(10),t(75)),s=function(){function e(e,n){for(var t=0;t<n.length;t++){var a=n[t];a.enumerable=a.enumerable||!1,a.configurable=!0,"value"in a&&(a.writable=!0),Object.defineProperty(e,a.key,a)}}return function(n,t,a){return t&&e(n.prototype,t),a&&e(n,a),n}}(),p=function e(n,t,a){null===n&&(n=Function.prototype);var i=Object.getOwnPropertyDescriptor(n,t);if(void 0===i){var o=Object.getPrototypeOf(n);return null===o?void 0:e(o,t,a)}if("value"in i)return i.value;var r=i.get;return void 0!==r?r.call(a):void 0},d=(a=["\n    <style include='ha-style'>\n      :host {\n        -ms-user-select: initial;\n        -webkit-user-select: initial;\n        -moz-user-select: initial;\n      }\n\n      .content {\n        padding: 16px;\n        max-width: 600px;\n        margin: 0 auto;\n      }\n\n      paper-card {\n        display: block;\n      }\n\n      paper-item {\n        cursor: pointer;\n      }\n\n      .empty {\n        text-align: center;\n        color: var(--secondary-text-color);\n      }\n\n      .header {\n        @apply --paper-font-title;\n      }\n\n      .row {\n        display: flex;\n       justify-content: space-between;\n      }\n      paper-dialog {\n        border-radius: 2px;\n      }\n      paper-dialog p {\n        color: var(--secondary-text-color);\n      }\n\n      #mp3dialog paper-icon-button {\n        float: right;\n      }\n\n      @media all and (max-width: 450px) {\n        paper-dialog {\n          margin: 0;\n          width: 100%;\n          max-height: calc(100% - 64px);\n\n          position: fixed !important;\n          bottom: 0px;\n          left: 0px;\n          right: 0px;\n          overflow: scroll;\n          border-bottom-left-radius: 0px;\n          border-bottom-right-radius: 0px;\n        }\n\n        .content {\n          width: auto;\n          padding: 0;\n        }\n      }\n\n      .tip {\n        color: var(--secondary-text-color);\n        font-size: 14px;\n      }\n      .date {\n        color: var(--primary-text-color);\n      }\n    </style>\n\n    <app-header-layout has-scrolling-region>\n      <app-header slot=\"header\" fixed>\n        <app-toolbar>\n          <ha-menu-button narrow='[[narrow]]' show-menu='[[showMenu]]'></ha-menu-button>\n          <div main-title>[[localize('panel.mailbox')]]</div>\n        </app-toolbar>\n      </app-header>\n      <div class='content'>\n        <paper-card>\n          <template is='dom-if' if='[[!_messages.length]]'>\n            <div class='card-content empty'>\n              [[localize('ui.panel.mailbox.empty')]]\n            </div>\n          </template>\n          <template is='dom-repeat' items='[[_messages]]'>\n            <paper-item on-click='openMP3Dialog'>\n              <paper-item-body style=\"width:100%\" two-line>\n                <div class=\"row\">\n                  <div>[[item.caller]]</div>\n                  <div class=\"tip\">[[localize('ui.duration.second', 'count', item.duration)]]</div>\n                </div>\n                <div secondary>\n                  <span class=\"date\">[[item.timestamp]]</span> - [[item.message]]\n                </div>\n              </paper-item-body>\n            </paper-item>\n          </template>\n        </paper-card>\n      </div>\n    </app-header-layout>\n\n    <paper-dialog with-backdrop id=\"mp3dialog\" on-iron-overlay-closed=\"_mp3Closed\">\n      <h2>\n        [[localize('ui.panel.mailbox.playback_title')]]\n        <paper-icon-button\n          on-click='openDeleteDialog'\n          icon='mdi:delete'\n        ></paper-icon-button>\n      </h2>\n      <div id=\"transcribe\"></div>\n      <div>\n        <audio id=\"mp3\" preload=\"none\" controls> <source id=\"mp3src\" src=\"\" type=\"audio/mpeg\" /></audio>\n      </div>\n    </paper-dialog>\n\n    <paper-dialog with-backdrop id=\"confirmdel\">\n      <p>[[localize('ui.panel.mailbox.delete_prompt')]]</p>\n      <div class=\"buttons\">\n        <paper-button dialog-dismiss>[[localize('ui.common.cancel')]]</paper-button>\n        <paper-button dialog-confirm autofocus on-click=\"deleteSelected\">[[localize('ui.panel.mailbox.delete_button')]]</paper-button>\n      </div>\n    </paper-dialog>\n    "],i=["\n    <style include='ha-style'>\n      :host {\n        -ms-user-select: initial;\n        -webkit-user-select: initial;\n        -moz-user-select: initial;\n      }\n\n      .content {\n        padding: 16px;\n        max-width: 600px;\n        margin: 0 auto;\n      }\n\n      paper-card {\n        display: block;\n      }\n\n      paper-item {\n        cursor: pointer;\n      }\n\n      .empty {\n        text-align: center;\n        color: var(--secondary-text-color);\n      }\n\n      .header {\n        @apply --paper-font-title;\n      }\n\n      .row {\n        display: flex;\n       justify-content: space-between;\n      }\n      paper-dialog {\n        border-radius: 2px;\n      }\n      paper-dialog p {\n        color: var(--secondary-text-color);\n      }\n\n      #mp3dialog paper-icon-button {\n        float: right;\n      }\n\n      @media all and (max-width: 450px) {\n        paper-dialog {\n          margin: 0;\n          width: 100%;\n          max-height: calc(100% - 64px);\n\n          position: fixed !important;\n          bottom: 0px;\n          left: 0px;\n          right: 0px;\n          overflow: scroll;\n          border-bottom-left-radius: 0px;\n          border-bottom-right-radius: 0px;\n        }\n\n        .content {\n          width: auto;\n          padding: 0;\n        }\n      }\n\n      .tip {\n        color: var(--secondary-text-color);\n        font-size: 14px;\n      }\n      .date {\n        color: var(--primary-text-color);\n      }\n    </style>\n\n    <app-header-layout has-scrolling-region>\n      <app-header slot=\"header\" fixed>\n        <app-toolbar>\n          <ha-menu-button narrow='[[narrow]]' show-menu='[[showMenu]]'></ha-menu-button>\n          <div main-title>[[localize('panel.mailbox')]]</div>\n        </app-toolbar>\n      </app-header>\n      <div class='content'>\n        <paper-card>\n          <template is='dom-if' if='[[!_messages.length]]'>\n            <div class='card-content empty'>\n              [[localize('ui.panel.mailbox.empty')]]\n            </div>\n          </template>\n          <template is='dom-repeat' items='[[_messages]]'>\n            <paper-item on-click='openMP3Dialog'>\n              <paper-item-body style=\"width:100%\" two-line>\n                <div class=\"row\">\n                  <div>[[item.caller]]</div>\n                  <div class=\"tip\">[[localize('ui.duration.second', 'count', item.duration)]]</div>\n                </div>\n                <div secondary>\n                  <span class=\"date\">[[item.timestamp]]</span> - [[item.message]]\n                </div>\n              </paper-item-body>\n            </paper-item>\n          </template>\n        </paper-card>\n      </div>\n    </app-header-layout>\n\n    <paper-dialog with-backdrop id=\"mp3dialog\" on-iron-overlay-closed=\"_mp3Closed\">\n      <h2>\n        [[localize('ui.panel.mailbox.playback_title')]]\n        <paper-icon-button\n          on-click='openDeleteDialog'\n          icon='mdi:delete'\n        ></paper-icon-button>\n      </h2>\n      <div id=\"transcribe\"></div>\n      <div>\n        <audio id=\"mp3\" preload=\"none\" controls> <source id=\"mp3src\" src=\"\" type=\"audio/mpeg\" /></audio>\n      </div>\n    </paper-dialog>\n\n    <paper-dialog with-backdrop id=\"confirmdel\">\n      <p>[[localize('ui.panel.mailbox.delete_prompt')]]</p>\n      <div class=\"buttons\">\n        <paper-button dialog-dismiss>[[localize('ui.common.cancel')]]</paper-button>\n        <paper-button dialog-confirm autofocus on-click=\"deleteSelected\">[[localize('ui.panel.mailbox.delete_button')]]</paper-button>\n      </div>\n    </paper-dialog>\n    "],Object.freeze(Object.defineProperties(a,{raw:{value:Object.freeze(i)}})));var c=function(e){function n(){return function(e,n){if(!(e instanceof n))throw new TypeError("Cannot call a class as a function")}(this,n),function(e,n){if(!e)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return!n||"object"!=typeof n&&"function"!=typeof n?e:n}(this,(n.__proto__||Object.getPrototypeOf(n)).apply(this,arguments))}return function(e,n){if("function"!=typeof n&&null!==n)throw new TypeError("Super expression must either be null or a function, not "+typeof n);e.prototype=Object.create(n&&n.prototype,{constructor:{value:e,enumerable:!1,writable:!0,configurable:!0}}),n&&(Object.setPrototypeOf?Object.setPrototypeOf(e,n):e.__proto__=n)}(n,window.hassMixins.LocalizeMixin(r["a"])),s(n,[{key:"connectedCallback",value:function(){p(n.prototype.__proto__||Object.getPrototypeOf(n.prototype),"connectedCallback",this).call(this),this.hassChanged=this.hassChanged.bind(this),this.hass.connection.subscribeEvents(this.hassChanged,"mailbox_updated").then(function(e){this._unsubEvents=e}.bind(this)),this.computePlatforms().then(function(e){this.platforms=e,this.hassChanged()}.bind(this))}},{key:"disconnectedCallback",value:function(){p(n.prototype.__proto__||Object.getPrototypeOf(n.prototype),"disconnectedCallback",this).call(this),this._unsubEvents&&this._unsubEvents()}},{key:"hassChanged",value:function(){this._messages||(this._messages=[]),this.getMessages().then(function(e){this._messages=e}.bind(this))}},{key:"openMP3Dialog",value:function(e){var n=e.model.item.platform;this.currentMessage=e.model.item,this.$.mp3dialog.open(),this.$.mp3src.src="/api/mailbox/media/"+n+"/"+e.model.item.sha,this.$.transcribe.innerText=e.model.item.message,this.$.mp3.load(),this.$.mp3.play()}},{key:"_mp3Closed",value:function(){this.$.mp3.pause()}},{key:"openDeleteDialog",value:function(){this.$.confirmdel.open()}},{key:"deleteSelected",value:function(){var e=this.currentMessage;this.hass.callApi("DELETE","mailbox/delete/"+e.platform+"/"+e.sha),this.$.mp3dialog.close()}},{key:"getMessages",value:function(){var e=this.platforms.map(function(e){return this.hass.callApi("GET","mailbox/messages/"+e).then(function(n){for(var t=[],a=n.length,i=0;i<a;i++){var o=Object(l.a)(new Date(1e3*n[i].info.origtime));t.push({timestamp:o,caller:n[i].info.callerid,message:n[i].text,sha:n[i].sha,duration:n[i].info.duration,platform:e})}return t})}.bind(this));return Promise.all(e).then(function(n){for(var t=e.length,a=[],i=0;i<t;i++)a=a.concat(n[i]);return a.sort(function(e,n){return new Date(n.timestamp)-new Date(e.timestamp)}),a})}},{key:"computePlatforms",value:function(){return this.hass.callApi("GET","mailbox/platforms")}}],[{key:"template",get:function(){return Object(o.a)(d)}},{key:"properties",get:function(){return{hass:{type:Object},narrow:{type:Boolean,value:!1},showMenu:{type:Boolean,value:!1},platforms:{type:Array},_messages:{type:Array},currentMessage:{type:Object}}}}]),n}();customElements.define("ha-panel-mailbox",c)}}]);