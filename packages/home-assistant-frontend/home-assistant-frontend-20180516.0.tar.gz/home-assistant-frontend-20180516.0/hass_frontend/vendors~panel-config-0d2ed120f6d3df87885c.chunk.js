(window.webpackJsonp=window.webpackJsonp||[]).push([[2],{172:function(e,t,n){"use strict";n.d(t,"b",function(){return i}),n.d(t,"a",function(){return E}),n.d(t,"c",function(){return B});var o={},r=[],a=[];function i(e,t){var n,i,l,p,s=a;for(p=arguments.length;p-- >2;)r.push(arguments[p]);for(t&&null!=t.children&&(r.length||r.push(t.children),delete t.children);r.length;)if((i=r.pop())&&void 0!==i.pop)for(p=i.length;p--;)r.push(i[p]);else"boolean"==typeof i&&(i=null),(l="function"!=typeof e)&&(null==i?i="":"number"==typeof i?i=String(i):"string"!=typeof i&&(l=!1)),l&&n?s[s.length-1]+=i:s===a?s=[i]:s.push(i),n=l;var d=new function(){};return d.nodeName=e,d.children=s,d.attributes=null==t?void 0:t,d.key=null==t?void 0:t.key,void 0!==o.vnode&&o.vnode(d),d}function l(e,t){for(var n in t)e[n]=t[n];return e}var p="function"==typeof Promise?Promise.resolve().then.bind(Promise.resolve()):setTimeout;var s=/acit|ex(?:s|g|n|p|$)|rph|ows|mnc|ntw|ine[ch]|zoo|^ord/i,d=[];function c(e){!e._dirty&&(e._dirty=!0)&&1==d.push(e)&&(o.debounceRendering||p)(u)}function u(){var e,t=d;for(d=[];e=t.pop();)e._dirty&&T(e)}function h(e,t){return e.normalizedNodeName===t||e.nodeName.toLowerCase()===t.toLowerCase()}function b(e){var t=l({},e.attributes);t.children=e.children;var n=e.nodeName.defaultProps;if(void 0!==n)for(var o in n)void 0===t[o]&&(t[o]=n[o]);return t}function f(e){var t=e.parentNode;t&&t.removeChild(e)}function v(e,t,n,o,r){if("className"===t&&(t="class"),"key"===t);else if("ref"===t)n&&n(null),o&&o(e);else if("class"!==t||r)if("style"===t){if(o&&"string"!=typeof o&&"string"!=typeof n||(e.style.cssText=o||""),o&&"object"==typeof o){if("string"!=typeof n)for(var a in n)a in o||(e.style[a]="");for(var a in o)e.style[a]="number"==typeof o[a]&&!1===s.test(a)?o[a]+"px":o[a]}}else if("dangerouslySetInnerHTML"===t)o&&(e.innerHTML=o.__html||"");else if("o"==t[0]&&"n"==t[1]){var i=t!==(t=t.replace(/Capture$/,""));t=t.toLowerCase().substring(2),o?n||e.addEventListener(t,m,i):e.removeEventListener(t,m,i),(e._listeners||(e._listeners={}))[t]=o}else if("list"!==t&&"type"!==t&&!r&&t in e)!function(e,t,n){try{e[t]=n}catch(e){}}(e,t,null==o?"":o),null!=o&&!1!==o||e.removeAttribute(t);else{var l=r&&t!==(t=t.replace(/^xlink\:?/,""));null==o||!1===o?l?e.removeAttributeNS("http://www.w3.org/1999/xlink",t.toLowerCase()):e.removeAttribute(t):"function"!=typeof o&&(l?e.setAttributeNS("http://www.w3.org/1999/xlink",t.toLowerCase(),o):e.setAttribute(t,o))}else e.className=o||""}function m(e){return this._listeners[e.type](o.event&&o.event(e)||e)}var y=[],g=0,_=!1,x=!1;function w(){for(var e;e=y.pop();)o.afterMount&&o.afterMount(e),e.componentDidMount&&e.componentDidMount()}function k(e,t,n,o,r,a){g++||(_=null!=r&&void 0!==r.ownerSVGElement,x=null!=e&&!("__preactattr_"in e));var i=C(e,t,n,o,a);return r&&i.parentNode!==r&&r.appendChild(i),--g||(x=!1,a||w()),i}function C(e,t,n,o,r){var a=e,i=_;if(null!=t&&"boolean"!=typeof t||(t=""),"string"==typeof t||"number"==typeof t)return e&&void 0!==e.splitText&&e.parentNode&&(!e._component||r)?e.nodeValue!=t&&(e.nodeValue=t):(a=document.createTextNode(t),e&&(e.parentNode&&e.parentNode.replaceChild(a,e),S(e,!0))),a.__preactattr_=!0,a;var l,p,s=t.nodeName;if("function"==typeof s)return function(e,t,n,o){var r=e&&e._component,a=r,i=e,l=r&&e._componentConstructor===t.nodeName,p=l,s=b(t);for(;r&&!p&&(r=r._parentComponent);)p=r.constructor===t.nodeName;r&&p&&(!o||r._component)?(O(r,s,3,n,o),e=r.base):(a&&!l&&(z(a),e=i=null),r=A(t.nodeName,s,n),e&&!r.nextBase&&(r.nextBase=e,i=null),O(r,s,1,n,o),e=r.base,i&&e!==i&&(i._component=null,S(i,!1)));return e}(e,t,n,o);if(_="svg"===s||"foreignObject"!==s&&_,s=String(s),(!e||!h(e,s))&&(l=s,(p=_?document.createElementNS("http://www.w3.org/2000/svg",l):document.createElement(l)).normalizedNodeName=l,a=p,e)){for(;e.firstChild;)a.appendChild(e.firstChild);e.parentNode&&e.parentNode.replaceChild(a,e),S(e,!0)}var d=a.firstChild,c=a.__preactattr_,u=t.children;if(null==c){c=a.__preactattr_={};for(var m=a.attributes,y=m.length;y--;)c[m[y].name]=m[y].value}return!x&&u&&1===u.length&&"string"==typeof u[0]&&null!=d&&void 0!==d.splitText&&null==d.nextSibling?d.nodeValue!=u[0]&&(d.nodeValue=u[0]):(u&&u.length||null!=d)&&function(e,t,n,o,r){var a,i,l,p,s,d=e.childNodes,c=[],u={},b=0,v=0,m=d.length,y=0,g=t?t.length:0;if(0!==m)for(var _=0;_<m;_++){var x=d[_],w=x.__preactattr_,k=g&&w?x._component?x._component.__key:w.key:null;null!=k?(b++,u[k]=x):(w||(void 0!==x.splitText?!r||x.nodeValue.trim():r))&&(c[y++]=x)}if(0!==g)for(var _=0;_<g;_++){p=t[_],s=null;var k=p.key;if(null!=k)b&&void 0!==u[k]&&(s=u[k],u[k]=void 0,b--);else if(!s&&v<y)for(a=v;a<y;a++)if(void 0!==c[a]&&(N=i=c[a],A=r,"string"==typeof(I=p)||"number"==typeof I?void 0!==N.splitText:"string"==typeof I.nodeName?!N._componentConstructor&&h(N,I.nodeName):A||N._componentConstructor===I.nodeName)){s=i,c[a]=void 0,a===y-1&&y--,a===v&&v++;break}s=C(s,p,n,o),l=d[_],s&&s!==e&&s!==l&&(null==l?e.appendChild(s):s===l.nextSibling?f(l):e.insertBefore(s,l))}var N,I,A;if(b)for(var _ in u)void 0!==u[_]&&S(u[_],!1);for(;v<=y;)void 0!==(s=c[y--])&&S(s,!1)}(a,u,n,o,x||null!=c.dangerouslySetInnerHTML),function(e,t,n){var o;for(o in n)t&&null!=t[o]||null==n[o]||v(e,o,n[o],n[o]=void 0,_);for(o in t)"children"===o||"innerHTML"===o||o in n&&t[o]===("value"===o||"checked"===o?e[o]:n[o])||v(e,o,n[o],n[o]=t[o],_)}(a,t.attributes,c),_=i,a}function S(e,t){var n=e._component;n?z(n):(null!=e.__preactattr_&&e.__preactattr_.ref&&e.__preactattr_.ref(null),!1!==t&&null!=e.__preactattr_||f(e),N(e))}function N(e){for(e=e.lastChild;e;){var t=e.previousSibling;S(e,!0),e=t}}var I={};function A(e,t,n){var o,r=I[e.name];if(e.prototype&&e.prototype.render?(o=new e(t,n),E.call(o,t,n)):((o=new E(t,n)).constructor=e,o.render=L),r)for(var a=r.length;a--;)if(r[a].constructor===e){o.nextBase=r[a].nextBase,r.splice(a,1);break}return o}function L(e,t,n){return this.constructor(e,n)}function O(e,t,n,r,a){e._disable||(e._disable=!0,(e.__ref=t.ref)&&delete t.ref,(e.__key=t.key)&&delete t.key,!e.base||a?e.componentWillMount&&e.componentWillMount():e.componentWillReceiveProps&&e.componentWillReceiveProps(t,r),r&&r!==e.context&&(e.prevContext||(e.prevContext=e.context),e.context=r),e.prevProps||(e.prevProps=e.props),e.props=t,e._disable=!1,0!==n&&(1!==n&&!1===o.syncComponentUpdates&&e.base?c(e):T(e,1,a)),e.__ref&&e.__ref(e))}function T(e,t,n,r){if(!e._disable){var a,i,p,s=e.props,d=e.state,c=e.context,u=e.prevProps||s,h=e.prevState||d,f=e.prevContext||c,v=e.base,m=e.nextBase,_=v||m,x=e._component,C=!1;if(v&&(e.props=u,e.state=h,e.context=f,2!==t&&e.shouldComponentUpdate&&!1===e.shouldComponentUpdate(s,d,c)?C=!0:e.componentWillUpdate&&e.componentWillUpdate(s,d,c),e.props=s,e.state=d,e.context=c),e.prevProps=e.prevState=e.prevContext=e.nextBase=null,e._dirty=!1,!C){a=e.render(s,d,c),e.getChildContext&&(c=l(l({},c),e.getChildContext()));var N,I,L=a&&a.nodeName;if("function"==typeof L){var E=b(a);(i=x)&&i.constructor===L&&E.key==i.__key?O(i,E,1,c,!1):(N=i,e._component=i=A(L,E,c),i.nextBase=i.nextBase||m,i._parentComponent=e,O(i,E,0,c,!1),T(i,1,n,!0)),I=i.base}else p=_,(N=x)&&(p=e._component=null),(_||1===t)&&(p&&(p._component=null),I=k(p,a,c,n||!v,_&&_.parentNode,!0));if(_&&I!==_&&i!==x){var B=_.parentNode;B&&I!==B&&(B.replaceChild(I,_),N||(_._component=null,S(_,!1)))}if(N&&z(N),e.base=I,I&&!r){for(var j=e,R=e;R=R._parentComponent;)(j=R).base=I;I._component=j,I._componentConstructor=j.constructor}}if(!v||n?y.unshift(e):C||(e.componentDidUpdate&&e.componentDidUpdate(u,h,f),o.afterUpdate&&o.afterUpdate(e)),null!=e._renderCallbacks)for(;e._renderCallbacks.length;)e._renderCallbacks.pop().call(e);g||r||w()}}function z(e){o.beforeUnmount&&o.beforeUnmount(e);var t=e.base;e._disable=!0,e.componentWillUnmount&&e.componentWillUnmount(),e.base=null;var n=e._component;n?z(n):t&&(t.__preactattr_&&t.__preactattr_.ref&&t.__preactattr_.ref(null),e.nextBase=t,f(t),function(e){var t=e.constructor.name;(I[t]||(I[t]=[])).push(e)}(e),N(t)),e.__ref&&e.__ref(null)}function E(e,t){this._dirty=!0,this.context=t,this.props=e,this.state=this.state||{}}function B(e,t,n){return k(n,e,{},!1,t,!1)}l(E.prototype,{setState:function(e,t){var n=this.state;this.prevState||(this.prevState=l({},n)),l(n,"function"==typeof e?e(n,this.props):e),t&&(this._renderCallbacks=this._renderCallbacks||[]).push(t),c(this)},forceUpdate:function(e){e&&(this._renderCallbacks=this._renderCallbacks||[]).push(e),T(this,2)},render:function(){}})},184:function(e,t,n){"use strict";n(3);var o=n(93),r=n(74),a=(n(137),n(4)),i=n(0);
/**
@license
Copyright (c) 2015 The Polymer Project Authors. All rights reserved.
This code may only be used under the BSD style license found at http://polymer.github.io/LICENSE.txt
The complete set of authors may be found at http://polymer.github.io/AUTHORS.txt
The complete set of contributors may be found at http://polymer.github.io/CONTRIBUTORS.txt
Code distributed by Google as part of the polymer project is also
subject to an additional IP rights grant found at http://polymer.github.io/PATENTS.txt
*/
Object(a.a)({_template:i["a"]`
    <style include="paper-dialog-shared-styles"></style>
    <slot></slot>
`,is:"paper-dialog",behaviors:[r.a,o.a],listeners:{"neon-animation-finish":"_onNeonAnimationFinish"},_renderOpened:function(){this.cancelAnimation(),this.playAnimation("entry")},_renderClosed:function(){this.cancelAnimation(),this.playAnimation("exit")},_onNeonAnimationFinish:function(){this.opened?this._finishRenderOpened():this._finishRenderClosed()}})},188:function(e,t,n){"use strict";n(3),n(20),n(39);var o=n(78),r=(n(72),n(47),n(23),n(4));
/**
@license
Copyright (c) 2015 The Polymer Project Authors. All rights reserved.
This code may only be used under the BSD style license found at http://polymer.github.io/LICENSE.txt
The complete set of authors may be found at http://polymer.github.io/AUTHORS.txt
The complete set of contributors may be found at http://polymer.github.io/CONTRIBUTORS.txt
Code distributed by Google as part of the polymer project is also
subject to an additional IP rights grant found at http://polymer.github.io/PATENTS.txt
*/
const a=document.createElement("template");a.setAttribute("style","display: none;"),a.innerHTML='<dom-module id="paper-fab">\n  <template strip-whitespace="">\n    <style include="paper-material-styles">\n      :host {\n        @apply --layout-vertical;\n        @apply --layout-center-center;\n\n        background: var(--paper-fab-background, var(--accent-color));\n        border-radius: 50%;\n        box-sizing: border-box;\n        color: var(--text-primary-color);\n        cursor: pointer;\n        height: 56px;\n        min-width: 0;\n        outline: none;\n        padding: 16px;\n        position: relative;\n        -moz-user-select: none;\n        -ms-user-select: none;\n        -webkit-user-select: none;\n        user-select: none;\n        width: 56px;\n        z-index: 0;\n\n        /* NOTE: Both values are needed, since some phones require the value `transparent`. */\n        -webkit-tap-highlight-color: rgba(0,0,0,0);\n        -webkit-tap-highlight-color: transparent;\n\n        @apply --paper-fab;\n      }\n\n      [hidden] {\n        display: none !important;\n      }\n\n      :host([mini]) {\n        width: 40px;\n        height: 40px;\n        padding: 8px;\n\n        @apply --paper-fab-mini;\n      }\n\n      :host([disabled]) {\n        color: var(--paper-fab-disabled-text, var(--paper-grey-500));\n        background: var(--paper-fab-disabled-background, var(--paper-grey-300));\n\n        @apply --paper-fab-disabled;\n      }\n\n      iron-icon {\n        @apply --paper-fab-iron-icon;\n      }\n\n      span {\n        width: 100%;\n        white-space: nowrap;\n        overflow: hidden;\n        text-overflow: ellipsis;\n        text-align: center;\n\n        @apply --paper-fab-label;\n      }\n\n      :host(.keyboard-focus) {\n        background: var(--paper-fab-keyboard-focus-background, var(--paper-pink-900));\n      }\n\n      :host([elevation="1"]) {\n        @apply --paper-material-elevation-1;\n      }\n\n      :host([elevation="2"]) {\n        @apply --paper-material-elevation-2;\n      }\n\n      :host([elevation="3"]) {\n        @apply --paper-material-elevation-3;\n      }\n\n      :host([elevation="4"]) {\n        @apply --paper-material-elevation-4;\n      }\n\n      :host([elevation="5"]) {\n        @apply --paper-material-elevation-5;\n      }\n    </style>\n\n    <iron-icon id="icon" hidden$="{{!_computeIsIconFab(icon, src)}}" src="[[src]]" icon="[[icon]]"></iron-icon>\n    <span hidden$="{{_computeIsIconFab(icon, src)}}">{{label}}</span>\n  </template>\n\n  \n</dom-module>',document.head.appendChild(a.content),Object(r.a)({is:"paper-fab",behaviors:[o.a],properties:{src:{type:String,value:""},icon:{type:String,value:""},mini:{type:Boolean,value:!1,reflectToAttribute:!0},label:{type:String,observer:"_labelChanged"}},_labelChanged:function(){this.setAttribute("aria-label",this.label)},_computeIsIconFab:function(e,t){return e.length>0||t.length>0}})},193:function(e,t,n){"use strict";n(3);var o=n(77),r=(n(23),n(20),n(4)),a=n(29);
/**
@license
Copyright (c) 2015 The Polymer Project Authors. All rights reserved.
This code may only be used under the BSD style license found at http://polymer.github.io/LICENSE.txt
The complete set of authors may be found at http://polymer.github.io/AUTHORS.txt
The complete set of contributors may be found at http://polymer.github.io/CONTRIBUTORS.txt
Code distributed by Google as part of the polymer project is also
subject to an additional IP rights grant found at http://polymer.github.io/PATENTS.txt
*/
const i=document.createElement("template");i.setAttribute("style","display: none;"),i.innerHTML='<dom-module id="paper-radio-button">\n  <template strip-whitespace="">\n    <style>\n      :host {\n        display: inline-block;\n        line-height: 0;\n        white-space: nowrap;\n        cursor: pointer;\n        @apply --paper-font-common-base;\n        --calculated-paper-radio-button-size: var(--paper-radio-button-size, 16px);\n        /* -1px is a sentinel for the default and is replace in `attached`. */\n        --calculated-paper-radio-button-ink-size: var(--paper-radio-button-ink-size, -1px);\n      }\n\n      :host(:focus) {\n        outline: none;\n      }\n\n      #radioContainer {\n        @apply --layout-inline;\n        @apply --layout-center-center;\n        position: relative;\n        width: var(--calculated-paper-radio-button-size);\n        height: var(--calculated-paper-radio-button-size);\n        vertical-align: middle;\n\n        @apply --paper-radio-button-radio-container;\n      }\n\n      #ink {\n        position: absolute;\n        top: 50%;\n        left: 50%;\n        right: auto;\n        width: var(--calculated-paper-radio-button-ink-size);\n        height: var(--calculated-paper-radio-button-ink-size);\n        color: var(--paper-radio-button-unchecked-ink-color, var(--primary-text-color));\n        opacity: 0.6;\n        pointer-events: none;\n        -webkit-transform: translate(-50%, -50%);\n        transform: translate(-50%, -50%);\n      }\n\n      #ink[checked] {\n        color: var(--paper-radio-button-checked-ink-color, var(--primary-color));\n      }\n\n      #offRadio, #onRadio {\n        position: absolute;\n        box-sizing: border-box;\n        top: 0;\n        left: 0;\n        width: 100%;\n        height: 100%;\n        border-radius: 50%;\n      }\n\n      #offRadio {\n        border: 2px solid var(--paper-radio-button-unchecked-color, var(--primary-text-color));\n        background-color: var(--paper-radio-button-unchecked-background-color, transparent);\n        transition: border-color 0.28s;\n      }\n\n      #onRadio {\n        background-color: var(--paper-radio-button-checked-color, var(--primary-color));\n        -webkit-transform: scale(0);\n        transform: scale(0);\n        transition: -webkit-transform ease 0.28s;\n        transition: transform ease 0.28s;\n        will-change: transform;\n      }\n\n      :host([checked]) #offRadio {\n        border-color: var(--paper-radio-button-checked-color, var(--primary-color));\n      }\n\n      :host([checked]) #onRadio {\n        -webkit-transform: scale(0.5);\n        transform: scale(0.5);\n      }\n\n      #radioLabel {\n        line-height: normal;\n        position: relative;\n        display: inline-block;\n        vertical-align: middle;\n        margin-left: var(--paper-radio-button-label-spacing, 10px);\n        white-space: normal;\n        color: var(--paper-radio-button-label-color, var(--primary-text-color));\n\n        @apply --paper-radio-button-label;\n      }\n\n      :host([checked]) #radioLabel {\n        @apply --paper-radio-button-label-checked;\n      }\n\n      #radioLabel:dir(rtl) {\n        margin-left: 0;\n        margin-right: var(--paper-radio-button-label-spacing, 10px);\n      }\n\n      #radioLabel[hidden] {\n        display: none;\n      }\n\n      /* disabled state */\n\n      :host([disabled]) #offRadio {\n        border-color: var(--paper-radio-button-unchecked-color, var(--primary-text-color));\n        opacity: 0.5;\n      }\n\n      :host([disabled][checked]) #onRadio {\n        background-color: var(--paper-radio-button-unchecked-color, var(--primary-text-color));\n        opacity: 0.5;\n      }\n\n      :host([disabled]) #radioLabel {\n        /* slightly darker than the button, so that it\'s readable */\n        opacity: 0.65;\n      }\n    </style>\n\n    <div id="radioContainer">\n      <div id="offRadio"></div>\n      <div id="onRadio"></div>\n    </div>\n\n    <div id="radioLabel"><slot></slot></div>\n  </template>\n\n  \n</dom-module>',document.head.appendChild(i.content),Object(r.a)({is:"paper-radio-button",behaviors:[o.a],hostAttributes:{role:"radio","aria-checked":!1,tabindex:0},properties:{ariaActiveAttribute:{type:String,value:"aria-checked"}},ready:function(){this._rippleContainer=this.$.radioContainer},attached:function(){Object(a.a)(this,function(){if("-1px"===this.getComputedStyleValue("--calculated-paper-radio-button-ink-size").trim()){var e=parseFloat(this.getComputedStyleValue("--calculated-paper-radio-button-size").trim()),t=Math.floor(3*e);t%2!=e%2&&t++,this.updateStyles({"--paper-radio-button-ink-size":t+"px"})}})}})},197:function(e,t,n){"use strict";n(3),n(9);var o=n(144),r=(n(193),n(4)),a=n(0),i=n(43);
/**
@license
Copyright (c) 2015 The Polymer Project Authors. All rights reserved.
This code may only be used under the BSD style license found at http://polymer.github.io/LICENSE.txt
The complete set of authors may be found at http://polymer.github.io/AUTHORS.txt
The complete set of contributors may be found at http://polymer.github.io/CONTRIBUTORS.txt
Code distributed by Google as part of the polymer project is also
subject to an additional IP rights grant found at http://polymer.github.io/PATENTS.txt
*/
Object(r.a)({_template:a["a"]`
    <style>
      :host {
        display: inline-block;
      }

      :host ::slotted(*) {
        padding: var(--paper-radio-group-item-padding, 12px);
      }
    </style>

    <slot></slot>
`,is:"paper-radio-group",behaviors:[o.a],hostAttributes:{role:"radiogroup"},properties:{attrForSelected:{type:String,value:"name"},selectedAttribute:{type:String,value:"checked"},selectable:{type:String,value:"paper-radio-button"},allowEmptySelection:{type:Boolean,value:!1}},select:function(e){var t=this._valueToItem(e);if(!t||!t.hasAttribute("disabled")){if(this.selected){var n=this._valueToItem(this.selected);if(this.selected==e){if(!this.allowEmptySelection)return void(n&&(n.checked=!0));e=""}n&&(n.checked=!1)}i.a.select.apply(this,[e]),this.fire("paper-radio-group-changed")}},_activateFocusedItem:function(){this._itemActivate(this._valueForItem(this.focusedItem),this.focusedItem)},_onUpKey:function(e){this._focusPrevious(),e.preventDefault(),this._activateFocusedItem()},_onDownKey:function(e){this._focusNext(),e.preventDefault(),this._activateFocusedItem()},_onLeftKey:function(e){o.b._onLeftKey.apply(this,arguments),this._activateFocusedItem()},_onRightKey:function(e){o.b._onRightKey.apply(this,arguments),this._activateFocusedItem()}})},198:function(e,t,n){"use strict";n(3),n(9);var o=n(21),r=n(11),a=n(35),i=(n(39),n(40)),l=n(30),p=(n(120),n(23),n(124),n(123),n(4)),s=n(0),d=n(2),c=n(14);
/**
@license
Copyright (c) 2016 The Polymer Project Authors. All rights reserved.
This code may only be used under the BSD style license found at http://polymer.github.io/LICENSE.txt
The complete set of authors may be found at http://polymer.github.io/AUTHORS.txt
The complete set of contributors may be found at http://polymer.github.io/CONTRIBUTORS.txt
Code distributed by Google as part of the polymer project is also
subject to an additional IP rights grant found at http://polymer.github.io/PATENTS.txt
*/
Object(p.a)({_template:s["a"]`
    <style include="paper-dropdown-menu-shared-styles">
      :host(:focus) {
        outline: none;
      }

      :host {
        width: 200px;  /* Default size of an <input> */
      }

      /**
       * All of these styles below are for styling the fake-input display
       */
      [slot="dropdown-trigger"] {
        box-sizing: border-box;
        position: relative;
        width: 100%;
        padding: 16px 0 8px 0;
      }

      :host([disabled]) [slot="dropdown-trigger"] {
        pointer-events: none;
        opacity: var(--paper-dropdown-menu-disabled-opacity, 0.33);
      }

      :host([no-label-float]) [slot="dropdown-trigger"] {
        padding-top: 8px;   /* If there's no label, we need less space up top. */
      }

      #input {
        @apply --paper-font-subhead;
        @apply --paper-font-common-nowrap;
        line-height: 1.5;
        border-bottom: 1px solid var(--paper-dropdown-menu-color, var(--secondary-text-color));
        color: var(--paper-dropdown-menu-color, var(--primary-text-color));
        width: 100%;
        box-sizing: border-box;
        padding: 12px 20px 0 0;   /* Right padding so that text doesn't overlap the icon */
        outline: none;
        @apply --paper-dropdown-menu-input;
      }

      #input:dir(rtl) {
        padding-right: 0px;
        padding-left: 20px;
      }

      :host([disabled]) #input {
        border-bottom: 1px dashed var(--paper-dropdown-menu-color, var(--secondary-text-color));
      }

      :host([invalid]) #input {
        border-bottom: 2px solid var(--paper-dropdown-error-color, var(--error-color));
      }

      :host([no-label-float]) #input {
        padding-top: 0;   /* If there's no label, we need less space up top. */
      }

      label {
        @apply --paper-font-subhead;
        @apply --paper-font-common-nowrap;
        display: block;
        position: absolute;
        bottom: 0;
        left: 0;
        right: 0;
        /**
         * The container has a 16px top padding, and there's 12px of padding
         * between the input and the label (from the input's padding-top)
         */
        top: 28px;
        box-sizing: border-box;
        width: 100%;
        padding-right: 20px;    /* Right padding so that text doesn't overlap the icon */
        text-align: left;
        transition-duration: .2s;
        transition-timing-function: cubic-bezier(.4,0,.2,1);
        color: var(--paper-dropdown-menu-color, var(--secondary-text-color));
        @apply --paper-dropdown-menu-label;
      }

      label:dir(rtl) {
        padding-right: 0px;
        padding-left: 20px;
      }

      :host([no-label-float]) label {
        top: 8px;
        /* Since the label doesn't need to float, remove the animation duration
        which slows down visibility changes (i.e. when a selection is made) */
        transition-duration: 0s;
      }

      label.label-is-floating {
        font-size: 12px;
        top: 8px;
      }

      label.label-is-hidden {
        visibility: hidden;
      }

      :host([focused]) label.label-is-floating {
        color: var(--paper-dropdown-menu-focus-color, var(--primary-color));
      }

      :host([invalid]) label.label-is-floating {
        color: var(--paper-dropdown-error-color, var(--error-color));
      }

      /**
       * Sets up the focused underline. It's initially hidden, and becomes
       * visible when it's focused.
       */
      label:after {
        background-color: var(--paper-dropdown-menu-focus-color, var(--primary-color));
        bottom: 7px;    /* The container has an 8px bottom padding */
        content: '';
        height: 2px;
        left: 45%;
        position: absolute;
        transition-duration: .2s;
        transition-timing-function: cubic-bezier(.4,0,.2,1);
        visibility: hidden;
        width: 8px;
        z-index: 10;
      }

      :host([invalid]) label:after {
        background-color: var(--paper-dropdown-error-color, var(--error-color));
      }

      :host([no-label-float]) label:after {
        bottom: 7px;    /* The container has a 8px bottom padding */
      }

      :host([focused]:not([disabled])) label:after {
        left: 0;
        visibility: visible;
        width: 100%;
      }

      iron-icon {
        position: absolute;
        right: 0px;
        bottom: 8px;    /* The container has an 8px bottom padding */
        @apply --paper-font-subhead;
        color: var(--disabled-text-color);
        @apply --paper-dropdown-menu-icon;
      }

      iron-icon:dir(rtl) {
        left: 0;
        right: auto;
      }

      :host([no-label-float]) iron-icon {
        margin-top: 0px;
      }

      .error {
        display: inline-block;
        visibility: hidden;
        color: var(--paper-dropdown-error-color, var(--error-color));
        @apply --paper-font-caption;
        position: absolute;
        left:0;
        right:0;
        bottom: -12px;
      }

      :host([invalid]) .error {
        visibility: visible;
      }
    </style>

    <!-- this div fulfills an a11y requirement for combobox, do not remove -->
    <span role="button"></span>
    <paper-menu-button id="menuButton" vertical-align="[[verticalAlign]]" horizontal-align="[[horizontalAlign]]" vertical-offset="[[_computeMenuVerticalOffset(noLabelFloat, verticalOffset)]]" disabled="[[disabled]]" no-animations="[[noAnimations]]" on-iron-select="_onIronSelect" on-iron-deselect="_onIronDeselect" opened="{{opened}}" close-on-activate="" allow-outside-scroll="[[allowOutsideScroll]]">
      <!-- support hybrid mode: user might be using paper-menu-button 1.x which distributes via <content> -->
      <div class="dropdown-trigger" slot="dropdown-trigger">
        <label class\$="[[_computeLabelClass(noLabelFloat,alwaysFloatLabel,hasContent)]]">
          [[label]]
        </label>
        <div id="input" tabindex="-1">&nbsp;</div>
        <iron-icon icon="paper-dropdown-menu:arrow-drop-down"></iron-icon>
        <span class="error">[[errorMessage]]</span>
      </div>
      <slot id="content" name="dropdown-content" slot="dropdown-content"></slot>
    </paper-menu-button>
`,is:"paper-dropdown-menu-light",behaviors:[o.a,r.a,l.a,a.a,i.a],properties:{selectedItemLabel:{type:String,notify:!0,readOnly:!0},selectedItem:{type:Object,notify:!0,readOnly:!0},value:{type:String,notify:!0,observer:"_valueChanged"},label:{type:String},placeholder:{type:String},opened:{type:Boolean,notify:!0,value:!1,observer:"_openedChanged"},allowOutsideScroll:{type:Boolean,value:!1},noLabelFloat:{type:Boolean,value:!1,reflectToAttribute:!0},alwaysFloatLabel:{type:Boolean,value:!1},noAnimations:{type:Boolean,value:!1},horizontalAlign:{type:String,value:"right"},verticalAlign:{type:String,value:"top"},verticalOffset:Number,hasContent:{type:Boolean,readOnly:!0}},listeners:{tap:"_onTap"},keyBindings:{"up down":"open",esc:"close"},hostAttributes:{tabindex:0,role:"combobox","aria-autocomplete":"none","aria-haspopup":"true"},observers:["_selectedItemChanged(selectedItem)"],attached:function(){var e=this.contentElement;e&&e.selectedItem&&this._setSelectedItem(e.selectedItem)},get contentElement(){for(var e=Object(d.b)(this.$.content).getDistributedNodes(),t=0,n=e.length;t<n;t++)if(e[t].nodeType===Node.ELEMENT_NODE)return e[t]},open:function(){this.$.menuButton.open()},close:function(){this.$.menuButton.close()},_onIronSelect:function(e){this._setSelectedItem(e.detail.item)},_onIronDeselect:function(e){this._setSelectedItem(null)},_onTap:function(e){c.findOriginalTarget(e)===this&&this.open()},_selectedItemChanged:function(e){var t="";t=e?e.label||e.getAttribute("label")||e.textContent.trim():"",this.value=t,this._setSelectedItemLabel(t)},_computeMenuVerticalOffset:function(e,t){return t||(e?-4:8)},_getValidity:function(e){return this.disabled||!this.required||this.required&&!!this.value},_openedChanged:function(){var e=this.opened?"true":"false",t=this.contentElement;t&&t.setAttribute("aria-expanded",e)},_computeLabelClass:function(e,t,n){var o="";return!0===e?n?"label-is-hidden":"":((n||!0===t)&&(o+=" label-is-floating"),o)},_valueChanged:function(){this.$.input&&this.$.input.textContent!==this.value&&(this.$.input.textContent=this.value),this._setHasContent(!!this.value)}})},332:function(e,t,n){"use strict";n(3);var o=n(4),r=n(2);
/**
@license
Copyright (c) 2015 The Polymer Project Authors. All rights reserved.
This code may only be used under the BSD style license found at http://polymer.github.io/LICENSE.txt
The complete set of authors may be found at http://polymer.github.io/AUTHORS.txt
The complete set of contributors may be found at http://polymer.github.io/CONTRIBUTORS.txt
Code distributed by Google as part of the polymer project is also
subject to an additional IP rights grant found at http://polymer.github.io/PATENTS.txt
*/
const a=Object(o.a)({is:"iron-label",listeners:{tap:"_tapHandler"},properties:{for:{type:String,value:"",reflectToAttribute:!0,observer:"_forChanged"},_forElement:Object},attached:function(){this._forChanged()},ready:function(){this._generateLabelId()},_generateLabelId:function(){if(!this.id){var e="iron-label-"+a._labelNumber++;Object(r.b)(this).setAttribute("id",e)}},_findTarget:function(){if(this.for){var e=Object(r.b)(this).getOwnerRoot();return Object(r.b)(e).querySelector("#"+this.for)}var t=Object(r.b)(this).querySelector("[iron-label-target]");return t||(t=Object(r.b)(this).firstElementChild),t},_tapHandler:function(e){this._forElement&&(Object(r.b)(e).localTarget!==this._forElement&&(this._forElement.focus(),this._forElement.click()))},_applyLabelledBy:function(){this._forElement&&Object(r.b)(this._forElement).setAttribute("aria-labelledby",this.id)},_forChanged:function(){this._forElement&&Object(r.b)(this._forElement).removeAttribute("aria-labelledby"),this._forElement=this._findTarget(),this._applyLabelledBy()}});a._labelNumber=0}}]);