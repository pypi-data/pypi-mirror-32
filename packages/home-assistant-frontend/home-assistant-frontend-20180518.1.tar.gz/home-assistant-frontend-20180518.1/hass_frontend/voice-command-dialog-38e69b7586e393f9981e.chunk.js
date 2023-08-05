/*! For license information please see voice-command-dialog-38e69b7586e393f9981e.chunk.js.LICENSE */
(window.webpackJsonp=window.webpackJsonp||[]).push([[24],{162:function(t,e,i){"use strict";i.d(e,"b",function(){return o}),i.d(e,"a",function(){return a}),i(2);var s=i(34),n=i(1);const o={hostAttributes:{role:"dialog",tabindex:"-1"},properties:{modal:{type:Boolean,value:!1},__readied:{type:Boolean,value:!1}},observers:["_modalChanged(modal, __readied)"],listeners:{tap:"_onDialogClick"},ready:function(){this.__prevNoCancelOnOutsideClick=this.noCancelOnOutsideClick,this.__prevNoCancelOnEscKey=this.noCancelOnEscKey,this.__prevWithBackdrop=this.withBackdrop,this.__readied=!0},_modalChanged:function(t,e){e&&(t?(this.__prevNoCancelOnOutsideClick=this.noCancelOnOutsideClick,this.__prevNoCancelOnEscKey=this.noCancelOnEscKey,this.__prevWithBackdrop=this.withBackdrop,this.noCancelOnOutsideClick=!0,this.noCancelOnEscKey=!0,this.withBackdrop=!0):(this.noCancelOnOutsideClick=this.noCancelOnOutsideClick&&this.__prevNoCancelOnOutsideClick,this.noCancelOnEscKey=this.noCancelOnEscKey&&this.__prevNoCancelOnEscKey,this.withBackdrop=this.withBackdrop&&this.__prevWithBackdrop))},_updateClosingReasonConfirmed:function(t){this.closingReason=this.closingReason||{},this.closingReason.confirmed=t},_onDialogClick:function(t){for(var e=Object(n.b)(t).path,i=0,s=e.indexOf(this);i<s;i++){var o=e[i];if(o.hasAttribute&&(o.hasAttribute("dialog-dismiss")||o.hasAttribute("dialog-confirm"))){this._updateClosingReasonConfirmed(o.hasAttribute("dialog-confirm")),this.close(),t.stopPropagation();break}}}},a=[s.a,o]},169:function(t,e,i){"use strict";i(2),i(19),i(26),i(43),i(72);const s=document.createElement("template");s.setAttribute("style","display: none;"),s.innerHTML='<dom-module id="paper-dialog-shared-styles">\n  <template>\n    <style>\n      :host {\n        display: block;\n        margin: 24px 40px;\n\n        background: var(--paper-dialog-background-color, var(--primary-background-color));\n        color: var(--paper-dialog-color, var(--primary-text-color));\n\n        @apply --paper-font-body1;\n        @apply --shadow-elevation-16dp;\n        @apply --paper-dialog;\n      }\n\n      :host > ::slotted(*) {\n        margin-top: 20px;\n        padding: 0 24px;\n      }\n\n      :host > ::slotted(.no-padding) {\n        padding: 0;\n      }\n\n      \n      :host > ::slotted(*:first-child) {\n        margin-top: 24px;\n      }\n\n      :host > ::slotted(*:last-child) {\n        margin-bottom: 24px;\n      }\n\n      /* In 1.x, this selector was `:host > ::content h2`. In 2.x <slot> allows\n      to select direct children only, which increases the weight of this\n      selector, so we have to re-define first-child/last-child margins below. */\n      :host > ::slotted(h2) {\n        position: relative;\n        margin: 0;\n\n        @apply --paper-font-title;\n        @apply --paper-dialog-title;\n      }\n\n      /* Apply mixin again, in case it sets margin-top. */\n      :host > ::slotted(h2:first-child) {\n        margin-top: 24px;\n        @apply --paper-dialog-title;\n      }\n\n      /* Apply mixin again, in case it sets margin-bottom. */\n      :host > ::slotted(h2:last-child) {\n        margin-bottom: 24px;\n        @apply --paper-dialog-title;\n      }\n\n      :host > ::slotted(.paper-dialog-buttons),\n      :host > ::slotted(.buttons) {\n        position: relative;\n        padding: 8px 8px 8px 24px;\n        margin: 0;\n\n        color: var(--paper-dialog-button-color, var(--primary-color));\n\n        @apply --layout-horizontal;\n        @apply --layout-end-justified;\n      }\n    </style>\n  </template>\n</dom-module>',document.head.appendChild(s.content)},189:function(t,e,i){"use strict";var s=i(6),n=i(162),o=i(45),a=i(11);e.a=Object(s.a)(t=>(class extends(Object(o.b)([a.a,n.a],t)){static get properties(){return{withBackdrop:{type:Boolean,value:!0}}}}))},398:function(t,e,i){"use strict";i.r(e),i(44),i(169),i(53);var s=i(0),n=i(3),o=i(189);customElements.define("ha-voice-command-dialog",class extends(Object(o.a)(n.a)){static get template(){return s["a"]`
    <style include="paper-dialog-shared-styles">
      iron-icon {
        margin-right: 8px;
      }

      .content {
        width: 450px;
        min-height: 80px;
        font-size: 18px;
        padding: 16px;
      }

      .messages {
        max-height: 50vh;
        overflow: auto;
      }

      .messages::after {
        content: "";
        clear: both;
        display: block;
      }

      .message {
        clear: both;
        margin: 8px 0;
        padding: 8px;
        border-radius: 15px;
      }

      .message.user {
        margin-left: 24px;
        float: right;
        text-align: right;
        border-bottom-right-radius: 0px;
        background-color: var(--light-primary-color);
        color: var(--primary-text-color);
      }

      .message.hass {
        margin-right: 24px;
        float: left;
        border-bottom-left-radius: 0px;
        background-color: var(--primary-color);
        color: var(--text-primary-color);
      }

      .message.error {
        background-color: var(--google-red-500);
        color: var(--text-primary-color);
      }

      .icon {
        text-align: center;
      }

      .icon paper-icon-button {
        height: 52px;
        width: 52px;
      }

      .interimTranscript {
        color: darkgrey;
      }

      [hidden] {
        display: none;
      }

      :host {
        border-radius: 2px;
      }

      @media all and (max-width: 450px) {
        :host {
          margin: 0;
          width: 100%;
          max-height: calc(100% - 64px);

          position: fixed !important;
          bottom: 0px;
          left: 0px;
          right: 0px;
          overflow: scroll;
          border-bottom-left-radius: 0px;
          border-bottom-right-radius: 0px;
        }

        .content {
          width: auto;
        }

        .messages {
          max-height: 68vh;
        }
      }
    </style>

    <div class="content">
      <div class="messages" id="messages">
        <template is="dom-repeat" items="[[_conversation]]" as="message">
          <div class\$="[[_computeMessageClasses(message)]]">[[message.text]]</div>
        </template>
      </div>
      <template is="dom-if" if="[[results]]">
        <div class="messages">
          <div class="message user">
            <span>{{results.final}}</span>
            <span class="interimTranscript">[[results.interim]]</span>
            …
          </div>
        </div>
      </template>
      <div class="icon" hidden\$="[[results]]">
        <paper-icon-button icon="mdi:text-to-speech" on-click="startListening"></paper-icon-button>
      </div>
    </div>
`}static get properties(){return{hass:Object,results:{type:Object,value:null,observer:"_scrollMessagesBottom"},_conversation:{type:Array,value:function(){return[{who:"hass",text:"How can I help?"}]},observer:"_scrollMessagesBottom"}}}static get observers(){return["dialogOpenChanged(opened)"]}initRecognition(){this.recognition=new webkitSpeechRecognition,this.recognition.onstart=function(){this.results={final:"",interim:""}}.bind(this),this.recognition.onerror=function(){this.recognition.abort();var t=this.results.final||this.results.interim;this.results=null,""===t&&(t="<Home Assistant did not hear anything>"),this.push("_conversation",{who:"user",text:t,error:!0})}.bind(this),this.recognition.onend=function(){if(null!=this.results){var t=this.results.final||this.results.interim;this.results=null,this.push("_conversation",{who:"user",text:t}),this.hass.callApi("post","conversation/process",{text:t}).then(function(t){this.push("_conversation",{who:"hass",text:t.speech.plain.speech})}.bind(this),function(){this.set(["_conversation",this._conversation.length-1,"error"],!0)}.bind(this))}}.bind(this),this.recognition.onresult=function(t){for(var e=this.results,i="",s="",n=t.resultIndex;n<t.results.length;n++)t.results[n].isFinal?i+=t.results[n][0].transcript:s+=t.results[n][0].transcript;this.results={interim:s,final:e.final+i}}.bind(this)}startListening(){this.recognition||this.initRecognition(),this.results={interim:"",final:""},this.recognition.start()}_scrollMessagesBottom(){setTimeout(()=>{this.$.messages.scrollTop=this.$.messages.scrollHeight,0!==this.$.messages.scrollTop&&this.$.dialog.fire("iron-resize")},10)}dialogOpenChanged(t){t?this.startListening():!t&&this.results&&this.recognition.abort()}_computeMessageClasses(t){return"message "+t.who+(t.error?" error":"")}})}}]);
//# sourceMappingURL=voice-command-dialog-38e69b7586e393f9981e.chunk.js.map