/*! For license information please see more-info-dialog-55df7f17d55de7700199.chunk.js.LICENSE */
(window.webpackJsonp=window.webpackJsonp||[]).push([[25],{132:function(t,e,i){"use strict";i(107);var a=i(20);i(2);const s=document.createElement("template");s.setAttribute("style","display: none;"),s.innerHTML='<dom-module id="ha-paper-slider">\n  <template strip-whitespace="">\n    <style include="paper-slider">\n      .pin > .slider-knob > .slider-knob-inner {\n        font-size:  var(--ha-paper-slider-pin-font-size, 10px);\n        line-height: normal;\n      }\n\n      .pin > .slider-knob > .slider-knob-inner::before {\n        top: unset;\n        margin-left: unset;\n\n        bottom: calc(15px + var(--calculated-paper-slider-height)/2);\n        left: 50%;\n        width: 2.6em;\n        height: 2.6em;\n\n        -webkit-transform-origin: left bottom;\n        transform-origin: left bottom;\n        -webkit-transform: rotate(-45deg) scale(0) translate(0);\n        transform: rotate(-45deg) scale(0) translate(0);\n      }\n\n      .pin.expand > .slider-knob > .slider-knob-inner::before {\n        -webkit-transform: rotate(-45deg) scale(1) translate(7px, -7px);\n        transform: rotate(-45deg) scale(1) translate(7px, -7px);\n      }\n\n      .pin > .slider-knob > .slider-knob-inner::after {\n        top: unset;\n        font-size: unset;\n\n        bottom: calc(15px + var(--calculated-paper-slider-height)/2);\n        left: 50%;\n        margin-left: -1.3em;\n        width: 2.6em;\n        height: 2.4em;\n\n        -webkit-transform-origin: center bottom;\n        transform-origin: center bottom;\n        -webkit-transform: scale(0) translate(0);\n        transform: scale(0) translate(0);\n      }\n\n      .pin.expand > .slider-knob > .slider-knob-inner::after {\n        -webkit-transform: scale(1) translate(0, -10px);\n        transform: scale(1) translate(0, -10px);\n      }\n    </style>\n  </template>\n\n\n</dom-module>',document.head.appendChild(s.content);{const t=customElements.get("paper-slider");let e;class i extends t{static get template(){if(!e){e=a.a.import(i.is,"template");const s=document.importNode(t.template.content,!0);s.querySelector("style").remove(),e.content.append(s)}return e}}customElements.define("ha-paper-slider",i)}},162:function(t,e,i){"use strict";i.d(e,"b",function(){return r}),i.d(e,"a",function(){return o}),i(2);var a=i(34),s=i(1);const r={hostAttributes:{role:"dialog",tabindex:"-1"},properties:{modal:{type:Boolean,value:!1},__readied:{type:Boolean,value:!1}},observers:["_modalChanged(modal, __readied)"],listeners:{tap:"_onDialogClick"},ready:function(){this.__prevNoCancelOnOutsideClick=this.noCancelOnOutsideClick,this.__prevNoCancelOnEscKey=this.noCancelOnEscKey,this.__prevWithBackdrop=this.withBackdrop,this.__readied=!0},_modalChanged:function(t,e){e&&(t?(this.__prevNoCancelOnOutsideClick=this.noCancelOnOutsideClick,this.__prevNoCancelOnEscKey=this.noCancelOnEscKey,this.__prevWithBackdrop=this.withBackdrop,this.noCancelOnOutsideClick=!0,this.noCancelOnEscKey=!0,this.withBackdrop=!0):(this.noCancelOnOutsideClick=this.noCancelOnOutsideClick&&this.__prevNoCancelOnOutsideClick,this.noCancelOnEscKey=this.noCancelOnEscKey&&this.__prevNoCancelOnEscKey,this.withBackdrop=this.withBackdrop&&this.__prevWithBackdrop))},_updateClosingReasonConfirmed:function(t){this.closingReason=this.closingReason||{},this.closingReason.confirmed=t},_onDialogClick:function(t){for(var e=Object(s.b)(t).path,i=0,a=e.indexOf(this);i<a;i++){var r=e[i];if(r.hasAttribute&&(r.hasAttribute("dialog-dismiss")||r.hasAttribute("dialog-confirm"))){this._updateClosingReasonConfirmed(r.hasAttribute("dialog-confirm")),this.close(),t.stopPropagation();break}}}},o=[a.a,r]},169:function(t,e,i){"use strict";i(2),i(19),i(26),i(43),i(72);const a=document.createElement("template");a.setAttribute("style","display: none;"),a.innerHTML='<dom-module id="paper-dialog-shared-styles">\n  <template>\n    <style>\n      :host {\n        display: block;\n        margin: 24px 40px;\n\n        background: var(--paper-dialog-background-color, var(--primary-background-color));\n        color: var(--paper-dialog-color, var(--primary-text-color));\n\n        @apply --paper-font-body1;\n        @apply --shadow-elevation-16dp;\n        @apply --paper-dialog;\n      }\n\n      :host > ::slotted(*) {\n        margin-top: 20px;\n        padding: 0 24px;\n      }\n\n      :host > ::slotted(.no-padding) {\n        padding: 0;\n      }\n\n      \n      :host > ::slotted(*:first-child) {\n        margin-top: 24px;\n      }\n\n      :host > ::slotted(*:last-child) {\n        margin-bottom: 24px;\n      }\n\n      /* In 1.x, this selector was `:host > ::content h2`. In 2.x <slot> allows\n      to select direct children only, which increases the weight of this\n      selector, so we have to re-define first-child/last-child margins below. */\n      :host > ::slotted(h2) {\n        position: relative;\n        margin: 0;\n\n        @apply --paper-font-title;\n        @apply --paper-dialog-title;\n      }\n\n      /* Apply mixin again, in case it sets margin-top. */\n      :host > ::slotted(h2:first-child) {\n        margin-top: 24px;\n        @apply --paper-dialog-title;\n      }\n\n      /* Apply mixin again, in case it sets margin-bottom. */\n      :host > ::slotted(h2:last-child) {\n        margin-bottom: 24px;\n        @apply --paper-dialog-title;\n      }\n\n      :host > ::slotted(.paper-dialog-buttons),\n      :host > ::slotted(.buttons) {\n        position: relative;\n        padding: 8px 8px 8px 24px;\n        margin: 0;\n\n        color: var(--paper-dialog-button-color, var(--primary-color));\n\n        @apply --layout-horizontal;\n        @apply --layout-end-justified;\n      }\n    </style>\n  </template>\n</dom-module>',document.head.appendChild(a.content)},175:function(t,e,i){"use strict";i(2),i(19);var a=i(162),s=(i(26),i(4)),r=i(0);Object(s.a)({_template:r["a"]`
    <style>

      :host {
        display: block;
        @apply --layout-relative;
      }

      :host(.is-scrolled:not(:first-child))::before {
        content: '';
        position: absolute;
        top: 0;
        left: 0;
        right: 0;
        height: 1px;
        background: var(--divider-color);
      }

      :host(.can-scroll:not(.scrolled-to-bottom):not(:last-child))::after {
        content: '';
        position: absolute;
        bottom: 0;
        left: 0;
        right: 0;
        height: 1px;
        background: var(--divider-color);
      }

      .scrollable {
        padding: 0 24px;

        @apply --layout-scroll;
        @apply --paper-dialog-scrollable;
      }

      .fit {
        @apply --layout-fit;
      }
    </style>

    <div id="scrollable" class="scrollable" on-scroll="updateScrollState">
      <slot></slot>
    </div>
`,is:"paper-dialog-scrollable",properties:{dialogElement:{type:Object}},get scrollTarget(){return this.$.scrollable},ready:function(){this._ensureTarget(),this.classList.add("no-padding")},attached:function(){this._ensureTarget(),requestAnimationFrame(this.updateScrollState.bind(this))},updateScrollState:function(){this.toggleClass("is-scrolled",this.scrollTarget.scrollTop>0),this.toggleClass("can-scroll",this.scrollTarget.offsetHeight<this.scrollTarget.scrollHeight),this.toggleClass("scrolled-to-bottom",this.scrollTarget.scrollTop+this.scrollTarget.offsetHeight>=this.scrollTarget.scrollHeight)},_ensureTarget:function(){this.dialogElement=this.dialogElement||this.parentElement,this.dialogElement&&this.dialogElement.behaviors&&this.dialogElement.behaviors.indexOf(a.b)>=0?(this.dialogElement.sizingTarget=this.scrollTarget,this.scrollTarget.classList.remove("fit")):this.dialogElement&&this.scrollTarget.classList.add("fit")}})},184:function(t,e){window.hassAttributeUtil=window.hassAttributeUtil||{},window.hassAttributeUtil.DOMAIN_DEVICE_CLASS={binary_sensor:["battery","cold","connectivity","door","garage_door","gas","heat","light","lock","moisture","motion","moving","occupancy","opening","plug","power","presence","problem","safety","smoke","sound","vibration","window"],cover:["garage"],sensor:["battery","humidity","illuminance","temperature"]},window.hassAttributeUtil.UNKNOWN_TYPE="json",window.hassAttributeUtil.ADD_TYPE="key-value",window.hassAttributeUtil.TYPE_TO_TAG={string:"ha-customize-string",json:"ha-customize-string",icon:"ha-customize-icon",boolean:"ha-customize-boolean",array:"ha-customize-array","key-value":"ha-customize-key-value"},window.hassAttributeUtil.LOGIC_STATE_ATTRIBUTES=window.hassAttributeUtil.LOGIC_STATE_ATTRIBUTES||{entity_picture:void 0,friendly_name:{type:"string",description:"Name"},icon:{type:"icon"},emulated_hue:{type:"boolean",domains:["emulated_hue"]},emulated_hue_name:{type:"string",domains:["emulated_hue"]},haaska_hidden:void 0,haaska_name:void 0,homebridge_hidden:{type:"boolean"},homebridge_name:{type:"string"},supported_features:void 0,attribution:void 0,custom_ui_more_info:{type:"string"},custom_ui_state_card:{type:"string"},device_class:{type:"array",options:window.hassAttributeUtil.DOMAIN_DEVICE_CLASS,description:"Device class",domains:["binary_sensor","cover","sensor"]},hidden:{type:"boolean",description:"Hide from UI"},assumed_state:{type:"boolean",domains:["switch","light","cover","climate","fan","group"]},initial_state:{type:"string",domains:["automation"]},unit_of_measurement:{type:"string"}}},189:function(t,e,i){"use strict";var a=i(6),s=i(162),r=i(45),o=i(11);e.a=Object(a.a)(t=>(class extends(Object(r.b)([o.a,s.a],t)){static get properties(){return{withBackdrop:{type:Boolean,value:!0}}}}))},396:function(t,e,i){"use strict";i.r(e),i(169),i(175);var a=i(0),s=i(3),r=(i(129),i(97),i(53),i(138),i(141),i(139),i(25),i(63),i(52),i(11));customElements.define("more-info-alarm_control_panel",class extends(Object(r.a)(s.a)){static get template(){return a["a"]`
    <style is="custom-style" include="iron-flex"></style>

    <div class="layout horizontal">
      <paper-input label="code" value="{{enteredCode}}" pattern="[[codeFormat]]" type="password" hidden\$="[[!codeFormat]]" disabled="[[!codeInputEnabled]]"></paper-input>
    </div>
    <div class="layout horizontal">
      <paper-button on-click="handleDisarmTap" hidden\$="[[!disarmButtonVisible]]" disabled="[[!codeValid]]">Disarm</paper-button>
      <paper-button on-click="handleHomeTap" hidden\$="[[!armHomeButtonVisible]]" disabled="[[!codeValid]]">Arm Home</paper-button>
      <paper-button on-click="handleAwayTap" hidden\$="[[!armAwayButtonVisible]]" disabled="[[!codeValid]]">Arm Away</paper-button>
    </div>
`}static get properties(){return{hass:{type:Object},stateObj:{type:Object,observer:"stateObjChanged"},enteredCode:{type:String,value:""},disarmButtonVisible:{type:Boolean,value:!1},armHomeButtonVisible:{type:Boolean,value:!1},armAwayButtonVisible:{type:Boolean,value:!1},codeInputEnabled:{type:Boolean,value:!1},codeFormat:{type:String,value:""},codeValid:{type:Boolean,computed:"validateCode(enteredCode, codeFormat)"}}}validateCode(t,e){var i=new RegExp(e);return null===e||i.test(t)}stateObjChanged(t,e){if(t){const e={codeFormat:t.attributes.code_format,armHomeButtonVisible:"disarmed"===t.state,armAwayButtonVisible:"disarmed"===t.state};e.disarmButtonVisible="armed_home"===t.state||"armed_away"===t.state||"armed_night"===t.state||"armed_custom_bypass"===t.state||"pending"===t.state||"triggered"===t.state,e.codeInputEnabled=e.disarmButtonVisible||"disarmed"===t.state,this.setProperties(e)}e&&setTimeout(()=>{this.fire("iron-resize")},500)}handleDisarmTap(){this.callService("alarm_disarm",{code:this.enteredCode})}handleHomeTap(){this.callService("alarm_arm_home",{code:this.enteredCode})}handleAwayTap(){this.callService("alarm_arm_away",{code:this.enteredCode})}callService(t,e){var i=e||{};i.entity_id=this.stateObj.entity_id,this.hass.callService("alarm_control_panel",t,i).then(function(){this.enteredCode=""}.bind(this))}}),i(136),customElements.define("more-info-automation",class extends s.a{static get template(){return a["a"]`
    <style>
      paper-button {
        color: var(--primary-color);
        font-weight: 500;
        top: 3px;
        height: 37px;
      }
    </style>

    <p>
      Last triggered:
      <ha-relative-time datetime="[[stateObj.attributes.last_triggered]]"></ha-relative-time>
    </p>

    <paper-button on-click="handleTriggerTapped">TRIGGER</paper-button>
`}static get properties(){return{hass:{type:Object},stateObj:{type:Object}}}handleTriggerTapped(){this.hass.callService("automation","trigger",{entity_id:this.stateObj.entity_id})}});var o=i(16);customElements.define("more-info-camera",class extends(Object(r.a)(s.a)){static get template(){return a["a"]`
    <style>
      :host {
        max-width:640px;
      }

      .camera-image {
        width: 100%;
      }
    </style>

    <img class="camera-image" src="[[computeCameraImageUrl(hass, stateObj, isVisible)]]" on-load="imageLoaded" alt="[[_computeStateName(stateObj)]]">
`}static get properties(){return{hass:{type:Object},stateObj:{type:Object},isVisible:{type:Boolean,value:!0}}}connectedCallback(){super.connectedCallback(),this.isVisible=!0}disconnectedCallback(){this.isVisible=!1,super.disconnectedCallback()}imageLoaded(){this.fire("iron-resize")}_computeStateName(t){return Object(o.a)(t)}computeCameraImageUrl(t,e,i){return t.demo?"/demo/webcam.jpg":e&&i?"/api/camera_proxy_stream/"+e.entity_id+"?token="+e.attributes.access_token:"data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7"}}),i(98),i(77),i(84),i(137);var n=i(7),l=i(17);customElements.define("ha-climate-control",class extends(Object(r.a)(s.a)){static get template(){return a["a"]`
    <style is="custom-style" include="iron-flex iron-flex-alignment"></style>
    <style>
      /* local DOM styles go here */
      :host {
        @apply --layout-flex;
        @apply --layout-horizontal;
        @apply --layout-justified;
      }
      .in-flux#target_temperature {
        color: var(--google-red-500);
      }
      #target_temperature {
        @apply --layout-self-center;
        font-size: 200%;
      }
      .control-buttons {
        font-size: 200%;
        text-align: right;
      }
      paper-icon-button {
        height: 48px;
        width: 48px;
      }
    </style>

    <!-- local DOM goes here -->
    <div id="target_temperature">
      [[value]] [[units]]
    </div>
    <div class="control-buttons">
      <div>
        <paper-icon-button icon="mdi:chevron-up" on-click="incrementValue"></paper-icon-button>
      </div>
      <div>
        <paper-icon-button icon="mdi:chevron-down" on-click="decrementValue"></paper-icon-button>
      </div>
    </div>
`}static get properties(){return{value:{type:Number,observer:"valueChanged"},units:{type:String},min:{type:Number},max:{type:Number},step:{type:Number,value:1}}}temperatureStateInFlux(t){this.$.target_temperature.classList.toggle("in-flux",t)}incrementValue(){const t=this.value+this.step;this.value<this.max&&(this.last_changed=Date.now(),this.temperatureStateInFlux(!0)),t<=this.max?t<=this.min?this.value=this.min:this.value=t:this.value=this.max}decrementValue(){const t=this.value-this.step;this.value>this.min&&(this.last_changed=Date.now(),this.temperatureStateInFlux(!0)),t>=this.min?this.value=t:this.value=this.min}valueChanged(){this.last_changed&&window.setTimeout(()=>{Date.now()-this.last_changed>=2e3&&(this.fire("change"),this.temperatureStateInFlux(!1),this.last_changed=null)},2010)}}),i(132);var d=i(86);function c(t,e){if(!t||!t.attributes.supported_features)return"";const i=t.attributes.supported_features;return Object.keys(e).map(t=>0!=(i&t)?e[t]:"").filter(t=>""!==t).join(" ")}customElements.define("more-info-climate",class extends(Object(r.a)(s.a)){static get template(){return a["a"]`
    <style is="custom-style" include="iron-flex"></style>
    <style>
      :host {
        color: var(--primary-text-color);
        --paper-input-container-input: {
          text-transform: capitalize;
        }
      }

      .container-on,
      .container-away_mode,
      .container-aux_heat,
      .container-temperature,
      .container-humidity,
      .container-operation_list,
      .container-fan_list,
      .container-swing_list {
        display: none;
      }

      .has-on .container-on,
      .has-away_mode .container-away_mode,
      .has-aux_heat .container-aux_heat,
      .has-target_temperature .container-temperature,
      .has-target_temperature_low .container-temperature,
      .has-target_temperature_high .container-temperature,
      .has-target_humidity .container-humidity,
      .has-operation_mode .container-operation_list,
      .has-fan_mode .container-fan_list,
      .has-swing_list .container-swing_list,
      .has-swing_mode .container-swing_list {
        display: block;
        margin-bottom: 5px;
      }

      .container-operation_list iron-icon,
      .container-fan_list iron-icon,
      .container-swing_list iron-icon {
        margin: 22px 16px 0 0;
      }

      paper-dropdown-menu {
        width: 100%;
      }

      paper-item {
        cursor: pointer;
      }

      ha-paper-slider {
        width: 100%;
      }

      .container-humidity .single-row {
          display: flex;
          height: 50px;
      }

      .target-humidity {
        width: 90px;
        font-size: 200%;
        margin: auto;
      }

      ha-climate-control.range-control-left,
      ha-climate-control.range-control-right {
        float: left;
        width: 46%;
      }
      ha-climate-control.range-control-left {
        margin-right: 4%;
      }
      ha-climate-control.range-control-right {
        margin-left: 4%;
      }

      .humidity {
        --paper-slider-active-color: var(--paper-blue-400);
        --paper-slider-secondary-color: var(--paper-blue-400);
      }

      .single-row {
        padding: 8px 0;
      }

      .capitalize {
        text-transform: capitalize;
      }
    </style>

    <div class\$="[[computeClassNames(stateObj)]]">

      <template is="dom-if" if="[[supportsOn(stateObj)]]">
        <div class="container-on">
          <div class="center horizontal layout single-row">
            <div class="flex">On / Off</div>
            <paper-toggle-button checked="[[onToggleChecked]]" on-change="onToggleChanged">
            </paper-toggle-button>
          </div>
        </div>
      </template>

      <div class="container-temperature">
        <div class\$="[[stateObj.attributes.operation_mode]]">
          <div hidden\$="[[!supportsTemperatureControls(stateObj)]]">Target
            Temperature</div>
          <template is="dom-if" if="[[supportsTemperature(stateObj)]]">
            <ha-climate-control value="[[stateObj.attributes.temperature]]" units="[[stateObj.attributes.unit_of_measurement]]" step="[[computeTemperatureStepSize(stateObj)]]" min="[[stateObj.attributes.min_temp]]" max="[[stateObj.attributes.max_temp]]" on-change="targetTemperatureChanged">
            </ha-climate-control>
          </template>
          <template is="dom-if" if="[[supportsTemperatureRange(stateObj)]]">
            <ha-climate-control value="[[stateObj.attributes.target_temp_low]]" units="[[stateObj.attributes.unit_of_measurement]]" step="[[computeTemperatureStepSize(stateObj)]]" min="[[stateObj.attributes.min_temp]]" max="[[stateObj.attributes.target_temp_high]]" class="range-control-left" on-change="targetTemperatureLowChanged">
            </ha-climate-control>
            <ha-climate-control value="[[stateObj.attributes.target_temp_high]]" units="[[stateObj.attributes.unit_of_measurement]]" step="[[computeTemperatureStepSize(stateObj)]]" min="[[stateObj.attributes.target_temp_low]]" max="[[stateObj.attributes.max_temp]]" class="range-control-right" on-change="targetTemperatureHighChanged">
            </ha-climate-control>
          </template>
        </div>
      </div>

      <template is="dom-if" if="[[supportsHumidity(stateObj)]]">
        <div class="container-humidity">
          <div>Target Humidity</div>
            <div class="single-row">
              <div class="target-humidity">[[stateObj.attributes.humidity]] %</div>
              <ha-paper-slider class="humidity" min="[[stateObj.attributes.min_humidity]]" max="[[stateObj.attributes.max_humidity]]" secondary-progress="[[stateObj.attributes.max_humidity]]" step="1" pin="" value="[[stateObj.attributes.humidity]]" on-change="targetHumiditySliderChanged" ignore-bar-touch="">
              </ha-paper-slider>
          </div>
        </div>
      </template>

      <template is="dom-if" if="[[supportsOperationMode(stateObj)]]">
        <div class="container-operation_list">
          <div class="controls">
            <paper-dropdown-menu class="capitalize" label-float="" dynamic-align="" label="Operation">
              <paper-listbox slot="dropdown-content" selected="{{operationIndex}}">
                <template is="dom-repeat" items="[[stateObj.attributes.operation_list]]" on-dom-change="handleOperationListUpdate">
                  <paper-item class="capitalize">[[item]]</paper-item>
                </template>
              </paper-listbox>
            </paper-dropdown-menu>
          </div>
        </div>
      </template>

      <template is="dom-if" if="[[supportsFanMode(stateObj)]]">
        <div class="container-fan_list">
          <paper-dropdown-menu label-float="" dynamic-align="" label="Fan Mode">
            <paper-listbox slot="dropdown-content" selected="{{fanIndex}}">
              <template is="dom-repeat" items="[[stateObj.attributes.fan_list]]" on-dom-change="handleFanListUpdate">
                <paper-item>[[item]]</paper-item>
              </template>
            </paper-listbox>
          </paper-dropdown-menu>
        </div>
      </template>

      <template is="dom-if" if="[[supportsSwingMode(stateObj)]]">
        <div class="container-swing_list">
          <paper-dropdown-menu label-float="" dynamic-align="" label="Swing Mode">
            <paper-listbox slot="dropdown-content" selected="{{swingIndex}}">
              <template is="dom-repeat" items="[[stateObj.attributes.swing_list]]" on-dom-change="handleSwingListUpdate">
                <paper-item>[[item]]</paper-item>
              </template>
            </paper-listbox>
          </paper-dropdown-menu>
        </div>
      </template>

      <template is="dom-if" if="[[supportsAwayMode(stateObj)]]">
        <div class="container-away_mode">
          <div class="center horizontal layout single-row">
            <div class="flex">Away Mode</div>
            <paper-toggle-button checked="[[awayToggleChecked]]" on-change="awayToggleChanged">
            </paper-toggle-button>
          </div>
        </div>
      </template>

      <template is="dom-if" if="[[supportsAuxHeat(stateObj)]]">
        <div class="container-aux_heat">
          <div class="center horizontal layout single-row">
            <div class="flex">Aux Heat</div>
            <paper-toggle-button checked="[[auxToggleChecked]]" on-change="auxToggleChanged">
            </paper-toggle-button>
          </div>
        </div>
      </template>
    </div>
`}static get properties(){return{hass:{type:Object},stateObj:{type:Object,observer:"stateObjChanged"},operationIndex:{type:Number,value:-1,observer:"handleOperationmodeChanged"},fanIndex:{type:Number,value:-1,observer:"handleFanmodeChanged"},swingIndex:{type:Number,value:-1,observer:"handleSwingmodeChanged"},awayToggleChecked:Boolean,auxToggleChecked:Boolean,onToggleChecked:Boolean}}stateObjChanged(t,e){t&&this.setProperties({awayToggleChecked:"on"===t.attributes.away_mode,auxToggleChecked:"on"===t.attributes.aux_heat,onToggleChecked:"off"!==t.state}),e&&(this._debouncer=l.a.debounce(this._debouncer,n.timeOut.after(500),()=>{this.fire("iron-resize")}))}handleOperationListUpdate(){this.operationIndex=-1,this.stateObj.attributes.operation_list&&(this.operationIndex=this.stateObj.attributes.operation_list.indexOf(this.stateObj.attributes.operation_mode))}handleSwingListUpdate(){this.swingIndex=-1,this.stateObj.attributes.swing_list&&(this.swingIndex=this.stateObj.attributes.swing_list.indexOf(this.stateObj.attributes.swing_mode))}handleFanListUpdate(){this.fanIndex=-1,this.stateObj.attributes.fan_list&&(this.fanIndex=this.stateObj.attributes.fan_list.indexOf(this.stateObj.attributes.fan_mode))}computeTemperatureStepSize(t){return t.attributes.target_temp_step?t.attributes.target_temp_step:-1!==t.attributes.unit_of_measurement.indexOf("F")?1:.5}supportsTemperatureControls(t){return this.supportsTemperature(t)||this.supportsTemperatureRange(t)}supportsTemperature(t){return 0!=(1&t.attributes.supported_features)}supportsTemperatureRange(t){return 0!=(6&t.attributes.supported_features)}supportsHumidity(t){return 0!=(8&t.attributes.supported_features)}supportsFanMode(t){return 0!=(64&t.attributes.supported_features)}supportsOperationMode(t){return 0!=(128&t.attributes.supported_features)}supportsSwingMode(t){return 0!=(512&t.attributes.supported_features)}supportsAwayMode(t){return 0!=(1024&t.attributes.supported_features)}supportsAuxHeat(t){return 0!=(2048&t.attributes.supported_features)}supportsOn(t){return 0!=(4096&t.attributes.supported_features)}computeClassNames(t){var e=[Object(d.a)(t,["current_temperature","current_humidity"]),c(t,{1:"has-target_temperature",2:"has-target_temperature_high",4:"has-target_temperature_low",8:"has-target_humidity",16:"has-target_humidity_high",32:"has-target_humidity_low",64:"has-fan_mode",128:"has-operation_mode",256:"has-hold_mode",512:"has-swing_mode",1024:"has-away_mode",2048:"has-aux_heat",4096:"has-on"})];return e.push("more-info-climate"),e.join(" ")}targetTemperatureChanged(t){const e=t.target.value;e!==this.stateObj.attributes.temperature&&this.callServiceHelper("set_temperature",{temperature:e})}targetTemperatureLowChanged(t){const e=t.currentTarget.value;e!==this.stateObj.attributes.target_temp_low&&this.callServiceHelper("set_temperature",{target_temp_low:e,target_temp_high:this.stateObj.attributes.target_temp_high})}targetTemperatureHighChanged(t){const e=t.currentTarget.value;e!==this.stateObj.attributes.target_temp_high&&this.callServiceHelper("set_temperature",{target_temp_low:this.stateObj.attributes.target_temp_low,target_temp_high:e})}targetHumiditySliderChanged(t){const e=t.target.value;e!==this.stateObj.attributes.humidity&&this.callServiceHelper("set_humidity",{humidity:e})}awayToggleChanged(t){const e="on"===this.stateObj.attributes.away_mode,i=t.target.checked;e!==i&&this.callServiceHelper("set_away_mode",{away_mode:i})}auxToggleChanged(t){const e="on"===this.stateObj.attributes.aux_heat,i=t.target.checked;e!==i&&this.callServiceHelper("set_aux_heat",{aux_heat:i})}onToggleChanged(t){const e="off"!==this.stateObj.state,i=t.target.checked;e!==i&&this.callServiceHelper(i?"turn_on":"turn_off",{})}handleFanmodeChanged(t){if(""===t||-1===t)return;const e=this.stateObj.attributes.fan_list[t];e!==this.stateObj.attributes.fan_mode&&this.callServiceHelper("set_fan_mode",{fan_mode:e})}handleOperationmodeChanged(t){if(""===t||-1===t)return;const e=this.stateObj.attributes.operation_list[t];e!==this.stateObj.attributes.operation_mode&&this.callServiceHelper("set_operation_mode",{operation_mode:e})}handleSwingmodeChanged(t){if(""===t||-1===t)return;const e=this.stateObj.attributes.swing_list[t];e!==this.stateObj.attributes.swing_mode&&this.callServiceHelper("set_swing_mode",{swing_mode:e})}callServiceHelper(t,e){e.entity_id=this.stateObj.entity_id,this.hass.callService("climate",t,e).then(()=>{this.stateObjChanged(this.stateObj)})}}),i(111),i(100),i(134),customElements.define("more-info-configurator",class extends s.a{static get template(){return a["a"]`
    <style is="custom-style" include="iron-flex"></style>
    <style>
      p {
        margin: 8px 0;
      }

      a {
        color: var(--primary-color);
      }

      p > img {
        max-width: 100%;
      }

      p.center {
        text-align: center;
      }

      p.error {
        color: #C62828;
      }

      p.submit {
        text-align: center;
        height: 41px;
      }

      paper-spinner {
        width: 14px;
        height: 14px;
        margin-right: 20px;
      }

      [hidden] {
        display: none;
      }
    </style>

    <div class="layout vertical">
      <template is="dom-if" if="[[isConfigurable]]">
        <ha-markdown content="[[stateObj.attributes.description]]"></ha-markdown>

        <p class="error" hidden\$="[[!stateObj.attributes.errors]]">
          [[stateObj.attributes.errors]]
        </p>

        <template is="dom-repeat" items="[[stateObj.attributes.fields]]">
          <paper-input label="[[item.name]]" name="[[item.id]]" type="[[item.type]]" on-change="fieldChanged"></paper-input>
        </template>

        <p class="submit" hidden\$="[[!stateObj.attributes.submit_caption]]">
          <paper-button raised="" disabled="[[isConfiguring]]" on-click="submitClicked">
            <paper-spinner active="[[isConfiguring]]" hidden="[[!isConfiguring]]" alt="Configuring"></paper-spinner>
            [[stateObj.attributes.submit_caption]]
          </paper-button>

        </p>

      </template>
    </div>
`}static get properties(){return{stateObj:{type:Object},action:{type:String,value:"display"},isConfigurable:{type:Boolean,computed:"computeIsConfigurable(stateObj)"},isConfiguring:{type:Boolean,value:!1},fieldInput:{type:Object,value:function(){return{}}}}}computeIsConfigurable(t){return"configure"===t.state}fieldChanged(t){var e=t.target;this.fieldInput[e.name]=e.value}submitClicked(){var t={configure_id:this.stateObj.attributes.configure_id,fields:this.fieldInput};this.isConfiguring=!0,this.hass.callService("configurator","configure",t).then(function(){this.isConfiguring=!1}.bind(this),function(){this.isConfiguring=!1}.bind(this))}}),i(147),i(122);{const t={128:"has-set_tilt_position"};class e extends s.a{static get template(){return a["a"]`
    <style is="custom-style" include="iron-flex"></style>
    <style>
      .current_position, .tilt {
        max-height: 0px;
        overflow: hidden;
      }
      .has-current_position .current_position,
      .has-set_tilt_position .tilt,
      .has-current_tilt_position .tilt
      {
        max-height: 90px;
      }

      [invisible] {
        visibility: hidden !important;
      }
    </style>
    <div class\$="[[computeClassNames(stateObj)]]">

      <div class="current_position">
        <div>Position</div>
        <ha-paper-slider min="0" max="100" value="{{coverPositionSliderValue}}" step="1" pin="" disabled="[[!entityObj.supportsSetPosition]]" on-change="coverPositionSliderChanged" ignore-bar-touch=""></ha-paper-slider>
      </div>

      <div class="tilt">
        <div>Tilt position</div>
        <div>
          <ha-cover-tilt-controls hidden\$="[[entityObj.isTiltOnly]]" hass="[[hass]]" state-obj="[[stateObj]]">
          </ha-cover-tilt-controls>
        </div>
        <ha-paper-slider min="0" max="100" value="{{coverTiltPositionSliderValue}}" step="1" pin="" disabled="[[!entityObj.supportsSetTiltPosition]]" on-change="coverTiltPositionSliderChanged" ignore-bar-touch=""></ha-paper-slider>
      </div>

    </div>
`}static get properties(){return{hass:{type:Object},stateObj:{type:Object,observer:"stateObjChanged"},entityObj:{type:Object,computed:"computeEntityObj(hass, stateObj)"},coverPositionSliderValue:{type:Number},coverTiltPositionSliderValue:{type:Number}}}computeEntityObj(t,e){return new window.CoverEntity(t,e)}stateObjChanged(t){t&&this.setProperties({coverPositionSliderValue:t.attributes.current_position,coverTiltPositionSliderValue:t.attributes.current_tilt_position})}computeClassNames(e){return[Object(d.a)(e,["current_position","current_tilt_position"]),c(e,t)].join(" ")}coverPositionSliderChanged(t){this.entityObj.setCoverPosition(t.target.value)}coverTiltPositionSliderChanged(t){this.entityObj.setCoverTiltPosition(t.target.value)}}customElements.define("more-info-cover",e)}i(184),customElements.define("ha-attributes",class extends s.a{static get template(){return a["a"]`
    <style is="custom-style" include="iron-flex iron-flex-alignment"></style>
    <style>
      .data-entry .value {
        max-width: 200px;
      }
      .attribution {
        color: var(--secondary-text-color);
        text-align: right;
      }
    </style>

    <div class="layout vertical">
      <template is="dom-repeat" items="[[computeDisplayAttributes(stateObj, filtersArray)]]" as="attribute">
        <div class="data-entry layout justified horizontal">
          <div class="key">[[formatAttribute(attribute)]]</div>
          <div class="value">[[formatAttributeValue(stateObj, attribute)]]</div>
        </div>
      </template>
      <div class="attribution" hidden\$="[[!computeAttribution(stateObj)]]">[[computeAttribution(stateObj)]]</div>
    </div>
`}static get properties(){return{stateObj:{type:Object},extraFilters:{type:String,value:""},filtersArray:{type:Array,computed:"computeFiltersArray(extraFilters)"}}}computeFiltersArray(t){return Object.keys(window.hassAttributeUtil.LOGIC_STATE_ATTRIBUTES).concat(t?t.split(","):[])}computeDisplayAttributes(t,e){return t?Object.keys(t.attributes).filter(function(t){return-1===e.indexOf(t)}):[]}formatAttribute(t){return t.replace(/_/g," ")}formatAttributeValue(t,e){var i=t.attributes[e];return null===i?"-":Array.isArray(i)?i.join(", "):i instanceof Object?JSON.stringify(i,null,2):i}computeAttribution(t){return t.attributes.attribution}}),customElements.define("more-info-default",class extends s.a{static get template(){return a["a"]`
    <ha-attributes state-obj="[[stateObj]]"></ha-attributes>
`}static get properties(){return{stateObj:{type:Object}}}}),customElements.define("more-info-fan",class extends(Object(r.a)(s.a)){static get template(){return a["a"]`
    <style is="custom-style" include="iron-flex"></style>
    <style>
      .container-speed_list,
      .container-direction,
      .container-oscillating {
        display: none;
      }

      .has-speed_list .container-speed_list,
      .has-direction .container-direction,
      .has-oscillating .container-oscillating {
        display: block;
      }

      paper-dropdown-menu {
        width: 100%;
      }

      paper-item {
        cursor: pointer;
      }
    </style>

    <div class\$="[[computeClassNames(stateObj)]]">

      <div class="container-speed_list">
        <paper-dropdown-menu label-float="" dynamic-align="" label="Speed">
          <paper-listbox slot="dropdown-content" selected="{{speedIndex}}">
            <template is="dom-repeat" items="[[stateObj.attributes.speed_list]]">
              <paper-item>[[item]]</paper-item>
            </template>
          </paper-listbox>
        </paper-dropdown-menu>
      </div>

      <div class="container-oscillating">
        <div class="center horizontal layout single-row">
          <div class="flex">Oscillate</div>
          <paper-toggle-button checked="[[oscillationToggleChecked]]" on-change="oscillationToggleChanged">
          </paper-toggle-button>
        </div>
      </div>

      <div class="container-direction">
        <div class="direction">
          <div>Direction</div>
          <paper-icon-button icon="mdi:rotate-left" on-click="onDirectionLeft" title="Left" disabled="[[computeIsRotatingLeft(stateObj)]]"></paper-icon-button>
          <paper-icon-button icon="mdi:rotate-right" on-click="onDirectionRight" title="Right" disabled="[[computeIsRotatingRight(stateObj)]]"></paper-icon-button>
        </div>
      </div>
    </div>

    <ha-attributes state-obj="[[stateObj]]" extra-filters="speed,speed_list,oscillating,direction"></ha-attributes>
`}static get properties(){return{hass:{type:Object},stateObj:{type:Object,observer:"stateObjChanged"},speedIndex:{type:Number,value:-1,observer:"speedChanged"},oscillationToggleChecked:{type:Boolean}}}stateObjChanged(t,e){t&&this.setProperties({oscillationToggleChecked:t.attributes.oscillating,speedIndex:t.attributes.speed_list?t.attributes.speed_list.indexOf(t.attributes.speed):-1}),e&&setTimeout(()=>{this.fire("iron-resize")},500)}computeClassNames(t){return"more-info-fan "+Object(d.a)(t,["oscillating","speed_list","direction"])}speedChanged(t){var e;""!==t&&-1!==t&&(e=this.stateObj.attributes.speed_list[t])!==this.stateObj.attributes.speed&&this.hass.callService("fan","turn_on",{entity_id:this.stateObj.entity_id,speed:e})}oscillationToggleChanged(t){var e=this.stateObj.attributes.oscillating,i=t.target.checked;e!==i&&this.hass.callService("fan","oscillate",{entity_id:this.stateObj.entity_id,oscillating:i})}onDirectionLeft(){this.hass.callService("fan","set_direction",{entity_id:this.stateObj.entity_id,direction:"reverse"})}onDirectionRight(){this.hass.callService("fan","set_direction",{entity_id:this.stateObj.entity_id,direction:"forward"})}computeIsRotatingLeft(t){return"reverse"===t.attributes.direction}computeIsRotatingRight(t){return"forward"===t.attributes.direction}});var p=i(1),u=i(15),h=i(73),m=i(120);customElements.define("more-info-group",class extends s.a{static get template(){return a["a"]`
    <style>
      .child-card {
        margin-bottom: 8px;
      }

      .child-card:last-child {
        margin-bottom: 0;
      }
    </style>

    <div id="groupedControlDetails"></div>
    <template is="dom-repeat" items="[[states]]" as="state">
      <div class="child-card">
        <state-card-content state-obj="[[state]]" hass="[[hass]]"></state-card-content>
      </div>
    </template>
`}static get properties(){return{hass:{type:Object},stateObj:{type:Object},states:{type:Array,computed:"computeStates(stateObj, hass)"}}}static get observers(){return["statesChanged(stateObj, states)"]}computeStates(t,e){for(var i=[],a=t.attributes.entity_id,s=0;s<a.length;s++){var r=e.states[a[s]];r&&i.push(r)}return i}statesChanged(t,e){let i=!1;if(e&&e.length>0){const a=e.find(t=>"on"===t.state)||e[0],s=Object(u.a)(a);if("group"!==s){i=Object.assign({},a,{entity_id:t.entity_id,attributes:Object.assign({},a.attributes)});for(let t=0;t<e.length;t++)if(s!==Object(u.a)(e[t])){i=!1;break}}}if(i)Object(h.a)(this.$.groupedControlDetails,"MORE-INFO-"+Object(m.a)(i).toUpperCase(),{stateObj:i,hass:this.hass});else{const t=Object(p.b)(this.$.groupedControlDetails);t.lastChild&&t.removeChild(t.lastChild)}}}),i(146),customElements.define("more-info-history_graph",class extends s.a{static get template(){return a["a"]`
    <style>
    :host {
      display: block;
      margin-bottom: 6px;
    }
    </style>
    <ha-history_graph-card hass="[[hass]]" state-obj="[[stateObj]]" in-dialog="">
    <ha-attributes state-obj="[[stateObj]]"></ha-attributes>
  </ha-history_graph-card>
`}static get properties(){return{hass:Object,stateObj:Object}}}),i(2),i(188),customElements.define("more-info-input_datetime",class extends s.a{static get template(){return a["a"]`
    <div class\$="[[computeClassNames(stateObj)]]">
      <template is="dom-if" if="[[doesHaveDate(stateObj)]]" restamp="">
        <div>
          <vaadin-date-picker id="dateInput" on-value-changed="dateTimeChanged" label="Date" value="{{selectedDate}}"></vaadin-date-picker>
        </div>
      </template>
      <template is="dom-if" if="[[doesHaveTime(stateObj)]]" restamp="">
        <div>
          <paper-input
            type='time'
            label='time'
            value='{{selectedTime}}'
          ></paper-input>
        </div>
      </template>
    </div>
`}constructor(){super(),this.is_ready=!1}static get properties(){return{hass:{type:Object},stateObj:{type:Object,observer:"stateObjChanged"},selectedDate:{type:String,observer:"dateTimeChanged"},selectedTime:{type:String,observer:"dateTimeChanged"}}}ready(){super.ready(),this.is_ready=!0}getDateString(t){return"unknown"===t.state?"":(e=t.attributes.month<10?"0":"",i=t.attributes.day<10?"0":"",t.attributes.year+"-"+e+t.attributes.month+"-"+i+t.attributes.day);var e,i}dateTimeChanged(){if(!this.is_ready)return;let t=!1;const e={entity_id:this.stateObj.entity_id};if(this.stateObj.attributes.has_time&&(t|=this.selectedTime!==`${this.stateObj.attributes.hour}:${this.stateObj.attributes.minute}`,e.time=this.selectedTime),this.stateObj.attributes.has_date){if(0===this.selectedDate.length)return;const i=new Date(this.selectedDate);t|=new Date(this.stateObj.attributes.year,this.stateObj.attributes.month-1,this.stateObj.attributes.day)!==i,e.date=this.selectedDate}t&&this.hass.callService("input_datetime","set_datetime",e)}stateObjChanged(t){this.is_ready=!1,t.attributes.has_time&&(this.selectedHour=t.attributes.hour,this.selectedMinute=t.attributes.minute),t.attributes.has_date&&(this.selectedDate=this.getDateString(t)),this.is_ready=!0}doesHaveDate(t){return t.attributes.has_date}doesHaveTime(t){return t.attributes.has_time}computeClassNames(t){return"more-info-input_datetime "+Object(d.a)(t,["has_time","has_date"])}}),customElements.define("ha-color-picker",class extends(Object(r.a)(s.a)){static get template(){return a["a"]`
    <style>
      :host {
        user-select: none;
        -webkit-user-select: none;
      }

      #canvas {
        position: relative;
        width: 100%;
        max-width: 330px;
      }
      #canvas > * {
        display: block;
      }
      #interactionLayer {
        color: white;
        position: absolute;
        cursor: crosshair;
        width: 100%;
        height: 100%;
        overflow: visible;
      }
      #backgroundLayer {
        width: 100%;
        overflow: visible;
        --wheel-bordercolor: var(--ha-color-picker-wheel-bordercolor, white);
        --wheel-borderwidth: var(--ha-color-picker-wheel-borderwidth, 3);
        --wheel-shadow: var(--ha-color-picker-wheel-shadow, rgb(15, 15, 15) 10px 5px 5px 0px);
      }

      #marker {
        fill: currentColor;
        stroke: var(--ha-color-picker-marker-bordercolor, white);
        stroke-width: var(--ha-color-picker-marker-borderwidth, 3);
        filter: url(#marker-shadow)
      }
      .dragging #marker {
      }

      #colorTooltip {
        display: none;
        fill: currentColor;
        stroke: var(--ha-color-picker-tooltip-bordercolor, white);
        stroke-width: var(--ha-color-picker-tooltip-borderwidth, 3);
      }

      .touch.dragging #colorTooltip {
        display: inherit;
      }

    </style>
    <div id="canvas">
      <svg id="interactionLayer">
        <defs>
          <filter id="marker-shadow" x="-50%" y="-50%" width="200%" height="200%" filterUnits="objectBoundingBox">
             <feOffset result="offOut" in="SourceAlpha" dx="2" dy="2"></feOffset>
             <feGaussianBlur result="blurOut" in="offOut" stdDeviation="2"></feGaussianBlur>
             <feComponentTransfer in="blurOut" result="alphaOut">
               <feFuncA type="linear" slope="0.3"></feFuncA>
             </feComponentTransfer>
             <feBlend in="SourceGraphic" in2="alphaOut" mode="normal"></feBlend>
          </filter>
        </defs>
      </svg>
      <canvas id="backgroundLayer"></canvas>
    </div>
`}static get properties(){return{hsColor:{type:Object},desiredHsColor:{type:Object,observer:"applyHsColor"},width:{type:Number,value:500},height:{type:Number,value:500},radius:{type:Number,value:225},hueSegments:{type:Number,value:0},saturationSegments:{type:Number,value:0},ignoreSegments:{type:Boolean,value:!1},throttle:{type:Number,value:500}}}ready(){super.ready(),this.setupLayers(),this.drawColorWheel(),this.drawMarker(),this.interactionLayer.addEventListener("mousedown",t=>this.onMouseDown(t)),this.interactionLayer.addEventListener("touchstart",t=>this.onTouchStart(t))}convertToCanvasCoordinates(t,e){var i=this.interactionLayer.createSVGPoint();i.x=t,i.y=e;var a=i.matrixTransform(this.interactionLayer.getScreenCTM().inverse());return{x:a.x,y:a.y}}onMouseDown(t){const e=this.convertToCanvasCoordinates(t.clientX,t.clientY);this.isInWheel(e.x,e.y)&&(this.onMouseSelect(t),this.canvas.classList.add("mouse","dragging"),this.addEventListener("mousemove",this.onMouseSelect),this.addEventListener("mouseup",this.onMouseUp))}onMouseUp(){this.canvas.classList.remove("mouse","dragging"),this.removeEventListener("mousemove",this.onMouseSelect)}onMouseSelect(t){requestAnimationFrame(()=>this.processUserSelect(t))}onTouchStart(t){var e=t.changedTouches[0];const i=this.convertToCanvasCoordinates(e.clientX,e.clientY);if(this.isInWheel(i.x,i.y)){if(t.target===this.marker)return t.preventDefault(),this.canvas.classList.add("touch","dragging"),this.addEventListener("touchmove",this.onTouchSelect),void this.addEventListener("touchend",this.onTouchEnd);this.tapBecameScroll=!1,this.addEventListener("touchend",this.onTap),this.addEventListener("touchmove",()=>{this.tapBecameScroll=!0},{passive:!0})}}onTap(t){this.tapBecameScroll||(t.preventDefault(),this.onTouchSelect(t))}onTouchEnd(){this.canvas.classList.remove("touch","dragging"),this.removeEventListener("touchmove",this.onTouchSelect)}onTouchSelect(t){requestAnimationFrame(()=>this.processUserSelect(t.changedTouches[0]))}processUserSelect(t){var e=this.convertToCanvasCoordinates(t.clientX,t.clientY),i=this.getColor(e.x,e.y);this.onColorSelect(i)}onColorSelect(t){if(this.setMarkerOnColor(t),this.ignoreSegments||(t=this.applySegmentFilter(t)),this.applyColorToCanvas(t),this.colorSelectIsThrottled)return clearTimeout(this.ensureFinalSelect),void(this.ensureFinalSelect=setTimeout(()=>{this.fireColorSelected(t)},this.throttle));this.fireColorSelected(t),this.colorSelectIsThrottled=!0,setTimeout(()=>{this.colorSelectIsThrottled=!1},this.throttle)}fireColorSelected(t){this.hsColor=t,this.fire("colorselected",{hs:{h:t.h,s:t.s}})}setMarkerOnColor(t){var e=t.s*this.radius,i=(t.h-180)/180*Math.PI,a=`translate(${-e*Math.cos(i)},${-e*Math.sin(i)})`;this.marker.setAttribute("transform",a),this.tooltip.setAttribute("transform",a)}applyColorToCanvas(t){this.interactionLayer.style.color=`hsl(${t.h}, 100%, ${100-50*t.s}%)`}applyHsColor(t){this.hsColor&&this.hsColor.h===t.h&&this.hsColor.s===t.s||(this.setMarkerOnColor(t),this.ignoreSegments||(t=this.applySegmentFilter(t)),this.hsColor=t,this.applyColorToCanvas(t))}getAngle(t,e){return Math.atan2(-e,-t)/Math.PI*180+180}isInWheel(t,e){return this.getDistance(t,e)<=1}getDistance(t,e){return Math.sqrt(t*t+e*e)/this.radius}getColor(t,e){var i=this.getAngle(t,e),a=this.getDistance(t,e);return{h:i,s:Math.min(a,1)}}applySegmentFilter(t){if(this.hueSegments){const e=360/this.hueSegments,i=e/2;t.h-=i,t.h<0&&(t.h+=360);const a=t.h%e;t.h-=a-e}if(this.saturationSegments)if(1===this.saturationSegments)t.s=1;else{var e=1/this.saturationSegments,i=1/(this.saturationSegments-1),a=Math.floor(t.s/e)*i;t.s=Math.min(a,1)}return t}setupLayers(){this.canvas=this.$.canvas,this.backgroundLayer=this.$.backgroundLayer,this.interactionLayer=this.$.interactionLayer,this.originX=this.width/2,this.originY=this.originX,this.backgroundLayer.width=this.width,this.backgroundLayer.height=this.height,this.interactionLayer.setAttribute("viewBox",`${-this.originX} ${-this.originY} ${this.width} ${this.height}`)}drawColorWheel(){let t,e,i,a;const s=this.backgroundLayer.getContext("2d"),r=this.originX,o=this.originY,n=this.radius,l=window.getComputedStyle(this.backgroundLayer,null),d=parseInt(l.getPropertyValue("--wheel-borderwidth"),10),c=l.getPropertyValue("--wheel-bordercolor").trim(),p=l.getPropertyValue("--wheel-shadow").trim();if("none"!==p){const s=p.split("px ");t=s.pop(),e=parseInt(s[0],10),i=parseInt(s[1],10),a=parseInt(s[2],10)||0}const u=n+d/2,h=n,m=n+d;"none"!==l.shadow&&(s.save(),s.beginPath(),s.arc(r,o,m,0,2*Math.PI,!1),s.shadowColor=t,s.shadowOffsetX=e,s.shadowOffsetY=i,s.shadowBlur=a,s.fillStyle="white",s.fill(),s.restore()),function(t,e){const i=360/(t=t||360),a=i/2;for(var n=0;n<=360;n+=i){var l=(n-a)*(Math.PI/180),d=(n+a+1)*(Math.PI/180);s.beginPath(),s.moveTo(r,o),s.arc(r,o,h,l,d,!1),s.closePath();var c=s.createRadialGradient(r,o,0,r,o,h);let t=100;if(c.addColorStop(0,`hsl(${n}, 100%, ${t}%)`),e>0){const i=1/e;let a=0;for(var p=1;p<e;p+=1){var u=t;t=100-50*(a=p*i),c.addColorStop(a,`hsl(${n}, 100%, ${u}%)`),c.addColorStop(a,`hsl(${n}, 100%, ${t}%)`)}c.addColorStop(a,`hsl(${n}, 100%, 50%)`)}c.addColorStop(1,`hsl(${n}, 100%, 50%)`),s.fillStyle=c,s.fill()}}(this.hueSegments,this.saturationSegments),d>0&&(s.beginPath(),s.arc(r,o,u,0,2*Math.PI,!1),s.lineWidth=d,s.strokeStyle=c,s.stroke())}drawMarker(){const t=this.interactionLayer,e=.08*this.radius,i=.15*this.radius,a=-3*i;t.marker=document.createElementNS("http://www.w3.org/2000/svg","circle"),t.marker.setAttribute("id","marker"),t.marker.setAttribute("r",e),this.marker=t.marker,t.appendChild(t.marker),t.tooltip=document.createElementNS("http://www.w3.org/2000/svg","circle"),t.tooltip.setAttribute("id","colorTooltip"),t.tooltip.setAttribute("r",i),t.tooltip.setAttribute("cx",0),t.tooltip.setAttribute("cy",a),this.tooltip=t.tooltip,t.appendChild(t.tooltip)}}),i(44),customElements.define("ha-labeled-slider",class extends s.a{static get template(){return a["a"]`
    <style>
      :host {
        display: block;
        padding-bottom: 16px;
      }

      .title {
        margin-bottom: 16px;
        opacity: var(--dark-primary-opacity);
      }

      iron-icon {
        float: left;
        margin-top: 4px;
        opacity: var(--dark-secondary-opacity);
      }

      .slider-container {
        margin-left: 24px;
      }

      ha-paper-slider {
        background-image: var(--ha-slider-background);
      }
    </style>

    <div class="title">[[caption]]</div>
    <iron-icon icon="[[icon]]"></iron-icon>
    <div class="slider-container">
      <ha-paper-slider min="[[min]]" max="[[max]]" value="{{value}}" ignore-bar-touch="[[ignoreBarTouch]]">
      </ha-paper-slider>
    </div>
`}static get properties(){return{caption:String,icon:String,min:Number,max:Number,ignoreBarTouch:Boolean,value:{type:Number,notify:!0}}}});{const t={1:"has-brightness",2:"has-color_temp",4:"has-effect_list",16:"has-color",128:"has-white_value"};class e extends(Object(r.a)(s.a)){static get template(){return a["a"]`
    <style is="custom-style" include="iron-flex"></style>
    <style>
      .effect_list {
        padding-bottom: 16px;
      }

      .effect_list, .brightness, .color_temp, .white_value {
        max-height: 0px;
        overflow: hidden;
        transition: max-height .5s ease-in;
      }

      .color_temp {
        --ha-slider-background: -webkit-linear-gradient(right, rgb(255, 160, 0) 0%, white 50%, rgb(166, 209, 255) 100%);
        /* The color temp minimum value shouldn't be rendered differently. It's not "off". */
        --paper-slider-knob-start-border-color: var(--primary-color);
      }

      ha-color-picker {
        display: block;
        width: 100%;

        max-height: 0px;
        overflow: hidden;
        transition: max-height .5s ease-in;
      }

      .has-effect_list.is-on .effect_list,
      .has-brightness .brightness,
      .has-color_temp.is-on .color_temp,
      .has-white_value.is-on .white_value {
        max-height: 84px;
      }

      .has-color.is-on ha-color-picker {
        max-height: 500px;
        overflow: visible;
        --ha-color-picker-wheel-borderwidth: 5;
        --ha-color-picker-wheel-bordercolor: white;
        --ha-color-picker-wheel-shadow: none;
        --ha-color-picker-marker-borderwidth: 2;
        --ha-color-picker-marker-bordercolor: white;
      }

      .is-unavailable .control {
        max-height: 0px;
      }

      paper-item {
        cursor: pointer;
      }
    </style>

    <div class\$="[[computeClassNames(stateObj)]]">

      <div class="control brightness">
        <ha-labeled-slider caption="Brightness" icon="mdi:brightness-5" max="255" value="{{brightnessSliderValue}}" on-change="brightnessSliderChanged" ignore-bar-touch=""></ha-labeled-slider>
      </div>

      <div class="control color_temp">
        <ha-labeled-slider caption="Color Temperature" icon="mdi:thermometer" min="[[stateObj.attributes.min_mireds]]" max="[[stateObj.attributes.max_mireds]]" value="{{ctSliderValue}}" on-change="ctSliderChanged" ignore-bar-touch=""></ha-labeled-slider>
      </div>

      <div class="control white_value">
        <ha-labeled-slider caption="White Value" icon="mdi:file-word-box" max="255" value="{{wvSliderValue}}" on-change="wvSliderChanged" ignore-bar-touch=""></ha-labeled-slider>
      </div>

      <ha-color-picker class="control color" on-colorselected="colorPicked" desired-hs-color="{{colorPickerColor}}" throttle="500" hue-segments="24" saturation-segments="8">
      </ha-color-picker>

      <div class="control effect_list">
        <paper-dropdown-menu label-float="" dynamic-align="" label="Effect">
          <paper-listbox slot="dropdown-content" selected="{{effectIndex}}">
            <template is="dom-repeat" items="[[stateObj.attributes.effect_list]]">
              <paper-item>[[item]]</paper-item>
            </template>
          </paper-listbox>
        </paper-dropdown-menu>
      </div>

      <ha-attributes state-obj="[[stateObj]]" extra-filters="brightness,color_temp,white_value,effect_list,effect,hs_color,rgb_color,xy_color,min_mireds,max_mireds"></ha-attributes>
    </div>
`}static get properties(){return{hass:{type:Object},stateObj:{type:Object,observer:"stateObjChanged"},effectIndex:{type:Number,value:-1,observer:"effectChanged"},brightnessSliderValue:{type:Number,value:0},ctSliderValue:{type:Number,value:0},wvSliderValue:{type:Number,value:0},colorPickerColor:{type:Object}}}stateObjChanged(t,e){const i={brightnessSliderValue:0};t&&"on"===t.state&&(i.brightnessSliderValue=t.attributes.brightness,i.ctSliderValue=t.attributes.color_temp,i.wvSliderValue=t.attributes.white_value,t.attributes.hs_color&&(i.colorPickerColor={h:t.attributes.hs_color[0],s:t.attributes.hs_color[1]/100}),t.attributes.effect_list?i.effectIndex=t.attributes.effect_list.indexOf(t.attributes.effect):i.effectIndex=-1),this.setProperties(i),e&&setTimeout(()=>{this.fire("iron-resize")},500)}computeClassNames(e){const i=[c(e,t)];return e&&"on"===e.state&&i.push("is-on"),e&&"unavailable"===e.state&&i.push("is-unavailable"),i.join(" ")}effectChanged(t){var e;""!==t&&-1!==t&&(e=this.stateObj.attributes.effect_list[t])!==this.stateObj.attributes.effect&&this.hass.callService("light","turn_on",{entity_id:this.stateObj.entity_id,effect:e})}brightnessSliderChanged(t){var e=parseInt(t.target.value,10);isNaN(e)||(0===e?this.hass.callService("light","turn_off",{entity_id:this.stateObj.entity_id}):this.hass.callService("light","turn_on",{entity_id:this.stateObj.entity_id,brightness:e}))}ctSliderChanged(t){var e=parseInt(t.target.value,10);isNaN(e)||this.hass.callService("light","turn_on",{entity_id:this.stateObj.entity_id,color_temp:e})}wvSliderChanged(t){var e=parseInt(t.target.value,10);isNaN(e)||this.hass.callService("light","turn_on",{entity_id:this.stateObj.entity_id,white_value:e})}serviceChangeColor(t,e,i){t.callService("light","turn_on",{entity_id:e,hs_color:[i.h,100*i.s]})}colorPicked(t){this.serviceChangeColor(this.hass,this.stateObj.entity_id,t.detail.hs)}}customElements.define("more-info-light",e)}customElements.define("more-info-lock",class extends s.a{static get template(){return a["a"]`
    <style>
      paper-input {
        display: inline-block;
      }
    </style>

    <div hidden\$="[[!stateObj.attributes.code_format]]">
      <paper-input label="code" value="{{enteredCode}}" pattern="[[stateObj.attributes.code_format]]" type="password"></paper-input>
      <paper-button on-click="handleUnlockTap" hidden\$="[[!isLocked]]">Unlock</paper-button>
      <paper-button on-click="handleLockTap" hidden\$="[[isLocked]]">Lock</paper-button>
    </div>
    <ha-attributes state-obj="[[stateObj]]" extra-filters="code_format"></ha-attributes>
`}static get properties(){return{hass:{type:Object},stateObj:{type:Object,observer:"stateObjChanged"},enteredCode:{type:String,value:""},isLocked:Boolean}}handleUnlockTap(){this.callService("unlock",{code:this.enteredCode})}handleLockTap(){this.callService("lock",{code:this.enteredCode})}stateObjChanged(t){t&&(this.isLocked="locked"===t.state)}callService(t,e){var i=e||{};i.entity_id=this.stateObj.entity_id,this.hass.callService("lock",t,i)}}),i(144);var b=i(118);{class t extends(Object(r.a)(s.a)){static get template(){return a["a"]`
    <style is="custom-style" include="iron-flex iron-flex-alignment"></style>
    <style>
      .media-state {
        text-transform: capitalize;
      }

      paper-icon-button[highlight] {
        color: var(--accent-color);
      }

      .volume {
        margin-bottom: 8px;

        max-height: 0px;
        overflow: hidden;
        transition: max-height .5s ease-in;
      }

      .has-volume_level .volume {
        max-height: 40px;
      }

      iron-icon.source-input {
        padding: 7px;
        margin-top: 15px;
      }

      paper-dropdown-menu.source-input {
        margin-left: 10px;
      }

      [hidden] {
        display: none !important;
      }

      paper-item {
        cursor: pointer;
      }
    </style>

    <div class\$="[[computeClassNames(stateObj)]]">
      <div class="layout horizontal">
        <div class="flex">
          <paper-icon-button icon="mdi:power" highlight\$="[[playerObj.isOff]]" on-click="handleTogglePower" hidden\$="[[computeHidePowerButton(playerObj)]]"></paper-icon-button>
        </div>
        <div>
          <template is="dom-if" if="[[computeShowPlaybackControls(playerObj)]]">
            <paper-icon-button icon="mdi:skip-previous" on-click="handlePrevious" hidden\$="[[!playerObj.supportsPreviousTrack]]"></paper-icon-button>
            <paper-icon-button icon="[[computePlaybackControlIcon(playerObj)]]" on-click="handlePlaybackControl" hidden\$="[[!computePlaybackControlIcon(playerObj)]]" highlight=""></paper-icon-button>
            <paper-icon-button icon="mdi:skip-next" on-click="handleNext" hidden\$="[[!playerObj.supportsNextTrack]]"></paper-icon-button>
          </template>
        </div>
      </div>
      <!-- VOLUME -->
      <div class="volume_buttons center horizontal layout" hidden\$="[[computeHideVolumeButtons(playerObj)]]">
        <paper-icon-button on-click="handleVolumeTap" icon="mdi:volume-off"></paper-icon-button>
        <paper-icon-button id="volumeDown" disabled\$="[[playerObj.isMuted]]" on-mousedown="handleVolumeDown" on-touchstart="handleVolumeDown" icon="mdi:volume-medium"></paper-icon-button>
        <paper-icon-button id="volumeUp" disabled\$="[[playerObj.isMuted]]" on-mousedown="handleVolumeUp" on-touchstart="handleVolumeUp" icon="mdi:volume-high"></paper-icon-button>
      </div>
      <div class="volume center horizontal layout" hidden\$="[[!playerObj.supportsVolumeSet]]">
        <paper-icon-button on-click="handleVolumeTap" hidden\$="[[playerObj.supportsVolumeButtons]]" icon="[[computeMuteVolumeIcon(playerObj)]]"></paper-icon-button>
        <ha-paper-slider disabled\$="[[playerObj.isMuted]]" min="0" max="100" value="[[playerObj.volumeSliderValue]]" on-change="volumeSliderChanged" class="flex" ignore-bar-touch="">
        </ha-paper-slider>
      </div>
      <!-- SOURCE PICKER -->
      <div class="controls layout horizontal justified" hidden\$="[[computeHideSelectSource(playerObj)]]">
        <iron-icon class="source-input" icon="mdi:login-variant"></iron-icon>
        <paper-dropdown-menu class="flex source-input" dynamic-align="" label-float="" label="Source">
          <paper-listbox slot="dropdown-content" selected="{{sourceIndex}}">
            <template is="dom-repeat" items="[[playerObj.sourceList]]">
              <paper-item>[[item]]</paper-item>
            </template>
          </paper-listbox>
        </paper-dropdown-menu>
      </div>
      <!-- TTS -->
      <div hidden\$="[[computeHideTTS(ttsLoaded, playerObj)]]" class="layout horizontal end">
        <paper-input id="ttsInput" label="Text to speak" class="flex" value="{{ttsMessage}}" on-keydown="ttsCheckForEnter"></paper-input>
        <paper-icon-button icon="mdi:send" on-click="sendTTS"></paper-icon-button>
      </div>
    </div>
`}static get properties(){return{hass:Object,stateObj:Object,playerObj:{type:Object,computed:"computePlayerObj(hass, stateObj)",observer:"playerObjChanged"},sourceIndex:{type:Number,value:0,observer:"handleSourceChanged"},ttsLoaded:{type:Boolean,computed:"computeTTSLoaded(hass)"},ttsMessage:{type:String,value:""}}}computePlayerObj(t,e){return new window.HassMediaPlayerEntity(t,e)}playerObjChanged(t,e){t&&void 0!==t.sourceList&&(this.sourceIndex=t.sourceList.indexOf(t.source)),e&&setTimeout(()=>{this.fire("iron-resize")},500)}computeClassNames(t){return Object(d.a)(t,["volume_level"])}computeMuteVolumeIcon(t){return t.isMuted?"mdi:volume-off":"mdi:volume-high"}computeHideVolumeButtons(t){return!t.supportsVolumeButtons||t.isOff}computeShowPlaybackControls(t){return!t.isOff&&t.hasMediaControl}computePlaybackControlIcon(t){return t.isPlaying?t.supportsPause?"mdi:pause":"mdi:stop":t.supportsPlay?"mdi:play":null}computeHidePowerButton(t){return t.isOff?!t.supportsTurnOn:!t.supportsTurnOff}computeHideSelectSource(t){return t.isOff||!t.supportsSelectSource||!t.sourceList}computeHideTTS(t,e){return!t||!e.supportsPlayMedia}computeTTSLoaded(t){return Object(b.a)(t,"tts")}handleTogglePower(){this.playerObj.togglePower()}handlePrevious(){this.playerObj.previousTrack()}handlePlaybackControl(){this.playerObj.mediaPlayPause()}handleNext(){this.playerObj.nextTrack()}handleSourceChanged(t,e){if(!this.playerObj||!this.playerObj.supportsSelectSource||void 0===this.playerObj.sourceList||t<0||t>=this.playerObj.sourceList||void 0===e)return;const i=this.playerObj.sourceList[t];i!==this.playerObj.source&&this.playerObj.selectSource(i)}handleVolumeTap(){this.playerObj.supportsVolumeMute&&this.playerObj.volumeMute(!this.playerObj.isMuted)}handleVolumeUp(){const t=this.$.volumeUp;this.handleVolumeWorker("volume_up",t,!0)}handleVolumeDown(){const t=this.$.volumeDown;this.handleVolumeWorker("volume_down",t,!0)}handleVolumeWorker(t,e,i){(i||void 0!==e&&e.pointerDown)&&(this.playerObj.callService(t),setTimeout(()=>this.handleVolumeWorker(t,e,!1),500))}volumeSliderChanged(t){const e=parseFloat(t.target.value),i=e>0?e/100:0;this.playerObj.setVolume(i)}ttsCheckForEnter(t){13===t.keyCode&&this.sendTTS()}sendTTS(){const t=this.hass.config.services.tts,e=Object.keys(t).sort();let i,a;for(a=0;a<e.length;a++)if(-1!==e[a].indexOf("_say")){i=e[a];break}i&&(this.hass.callService("tts",i,{entity_id:this.stateObj.entity_id,message:this.ttsMessage}),this.ttsMessage="",this.$.ttsInput.focus())}}customElements.define("more-info-media_player",t)}customElements.define("more-info-script",class extends s.a{static get template(){return a["a"]`
    <style is="custom-style" include="iron-flex iron-flex-alignment"></style>

    <div class="layout vertical">
      <div class="data-entry layout justified horizontal">
        <div class="key">Last Action</div>
        <div class="value">[[stateObj.attributes.last_action]]</div>
      </div>
    </div>
`}static get properties(){return{stateObj:{type:Object}}}});var g=i(87);customElements.define("more-info-sun",class extends s.a{static get template(){return a["a"]`
    <style is="custom-style" include="iron-flex iron-flex-alignment"></style>

    <template is="dom-repeat" items="[[computeOrder(risingDate, settingDate)]]">
      <div class="data-entry layout justified horizontal">
        <div class="key">
          <span>[[itemCaption(item)]]</span>
          <ha-relative-time datetime-obj="[[itemDate(item)]]"></ha-relative-time>
        </div>
        <div class="value">[[itemValue(item)]]</div>
      </div>
    </template>
      <div class="data-entry layout justified horizontal">
        <div class="key">Elevation</div>
        <div class="value">[[stateObj.attributes.elevation]]</div>
     </div>
`}static get properties(){return{stateObj:{type:Object},risingDate:{type:Object,computed:"computeRising(stateObj)"},settingDate:{type:Object,computed:"computeSetting(stateObj)"}}}computeRising(t){return new Date(t.attributes.next_rising)}computeSetting(t){return new Date(t.attributes.next_setting)}computeOrder(t,e){return t>e?["set","ris"]:["ris","set"]}itemCaption(t){return"ris"===t?"Rising ":"Setting "}itemDate(t){return"ris"===t?this.risingDate:this.settingDate}itemValue(t){return Object(g.a)(this.itemDate(t))}}),customElements.define("more-info-updater",class extends s.a{static get template(){return a["a"]`
    <style>
      .link {
        color: #03A9F4;
      }
    </style>

    <div>
      <a class="link" href="https://www.home-assistant.io/docs/installation/updating/" target="_blank">Update Instructions</a>
    </div>
`}static get properties(){return{stateObj:{type:Object}}}computeReleaseNotes(t){return t.attributes.release_notes||"https://www.home-assistant.io/docs/installation/updating/"}}),customElements.define("more-info-vacuum",class extends s.a{static get template(){return a["a"]`
    <style is="custom-style" include="iron-flex iron-flex-alignment"></style>
    <style>
      :host {
        @apply --paper-font-body1;
        line-height: 1.5;
      }

      .status-subtitle {
        color: var(--secondary-text-color);
      }

      paper-item {
        cursor: pointer;
      }

    </style>

    <div class="horizontal justified layout">
      <div hidden\$="[[!supportsStatus(stateObj)]]">
        <span class="status-subtitle">Status: </span><span><strong>[[stateObj.attributes.status]]</strong></span>
      </div>
      <div hidden\$="[[!supportsBattery(stateObj)]]">
        <span><iron-icon icon="[[stateObj.attributes.battery_icon]]"></iron-icon> [[stateObj.attributes.battery_level]] %</span>
      </div>
    </div>
    <div hidden\$="[[!supportsCommandBar(stateObj)]]">
      <p></p>
      <div class="status-subtitle">Vacuum cleaner commands:</div>
      <div class="horizontal justified layout">
        <div hidden\$="[[!supportsPause(stateObj)]]">
          <paper-icon-button icon="mdi:play-pause" on-click="onPlayPause" title="Start/Pause"></paper-icon-button>
        </div>
        <div hidden\$="[[!supportsStop(stateObj)]]">
          <paper-icon-button icon="mdi:stop" on-click="onStop" title="Stop"></paper-icon-button>
        </div>
        <div hidden\$="[[!supportsCleanSpot(stateObj)]]">
        <paper-icon-button icon="mdi:broom" on-click="onCleanSpot" title="Clean spot"></paper-icon-button>
        </div>
        <div hidden\$="[[!supportsLocate(stateObj)]]">
        <paper-icon-button icon="mdi:map-marker" on-click="onLocate" title="Locate"></paper-icon-button>
        </div>
        <div hidden\$="[[!supportsReturnHome(stateObj)]]">
        <paper-icon-button icon="mdi:home-map-marker" on-click="onReturnHome" title="Return home"></paper-icon-button>
        </div>
      </div>
    </div>

    <div hidden\$="[[!supportsFanSpeed(stateObj)]]">
      <div class="horizontal justified layout">
        <paper-dropdown-menu label-float="" dynamic-align="" label="Fan speed">
          <paper-listbox slot="dropdown-content" selected="{{fanSpeedIndex}}">
            <template is="dom-repeat" items="[[stateObj.attributes.fan_speed_list]]">
              <paper-item>[[item]]</paper-item>
            </template>
          </paper-listbox>
        </paper-dropdown-menu>
        <div style="justify-content: center; align-self: center; padding-top: 1.3em">
          <span><iron-icon icon="mdi:fan"></iron-icon> [[stateObj.attributes.fan_speed]]</span>
        </div>
      </div>
      <p></p>
    </div>
    <ha-attributes state-obj="[[stateObj]]" extra-filters="fan_speed,fan_speed_list,status,battery_level,battery_icon"></ha-attributes>
`}static get properties(){return{hass:{type:Object},inDialog:{type:Boolean,value:!1},stateObj:{type:Object},fanSpeedIndex:{type:Number,value:-1,observer:"fanSpeedChanged"}}}supportsPause(t){return 0!=(4&t.attributes.supported_features)}supportsStop(t){return 0!=(8&t.attributes.supported_features)}supportsReturnHome(t){return 0!=(16&t.attributes.supported_features)}supportsFanSpeed(t){return 0!=(32&t.attributes.supported_features)}supportsBattery(t){return 0!=(64&t.attributes.supported_features)}supportsStatus(t){return 0!=(128&t.attributes.supported_features)}supportsLocate(t){return 0!=(512&t.attributes.supported_features)}supportsCleanSpot(t){return 0!=(1024&t.attributes.supported_features)}supportsCommandBar(t){return 0!=(4&t.attributes.supported_features)|0!=(8&t.attributes.supported_features)|0!=(16&t.attributes.supported_features)|0!=(512&t.attributes.supported_features)|0!=(1024&t.attributes.supported_features)}fanSpeedChanged(t){var e;""!==t&&-1!==t&&(e=this.stateObj.attributes.fan_speed_list[t])!==this.stateObj.attributes.fan_speed&&this.hass.callService("vacuum","set_fan_speed",{entity_id:this.stateObj.entity_id,fan_speed:e})}onStop(){this.hass.callService("vacuum","stop",{entity_id:this.stateObj.entity_id})}onPlayPause(){this.hass.callService("vacuum","start_pause",{entity_id:this.stateObj.entity_id})}onLocate(){this.hass.callService("vacuum","locate",{entity_id:this.stateObj.entity_id})}onCleanSpot(){this.hass.callService("vacuum","clean_spot",{entity_id:this.stateObj.entity_id})}onReturnHome(){this.hass.callService("vacuum","return_to_base",{entity_id:this.stateObj.entity_id})}});var y=i(14);customElements.define("more-info-weather",class extends(Object(y.a)(s.a)){static get template(){return a["a"]`
    <style>
      iron-icon {
        color: var(--paper-item-icon-color);
      }
      .section {
        margin: 16px 0 8px 0;
        font-size: 1.2em;
      }

      .flex {
        display: flex;
        height: 32px;
        align-items: center;
      }

      .main {
        flex: 1;
        margin-left: 24px;
      }

      .temp,
      .templow {
        min-width: 48px;
        text-align: right;
      }

      .templow {
        margin: 0 16px;
        color: var(--secondary-text-color);
      }

      .attribution {
        color: var(--secondary-text-color);
        text-align: center;
      }
    </style>

    <div class="flex">
      <iron-icon icon="mdi:thermometer"></iron-icon>
      <div class="main">[[localize('ui.card.weather.attributes.temperature')]]</div>
      <div>[[stateObj.attributes.temperature]] [[getUnit('temperature')]]</div>
    </div>
    <template is="dom-if" if="[[stateObj.attributes.pressure]]">
      <div class="flex">
        <iron-icon icon="mdi:gauge"></iron-icon>
        <div class="main">[[localize('ui.card.weather.attributes.air_pressure')]]</div>
        <div>[[stateObj.attributes.pressure]] hPa</div>
      </div>
    </template>
    <template is="dom-if" if="[[stateObj.attributes.humidity]]">
      <div class="flex">
        <iron-icon icon="mdi:water-percent"></iron-icon>
        <div class="main">[[localize('ui.card.weather.attributes.humidity')]]</div>
        <div>[[stateObj.attributes.humidity]] %</div>
      </div>
    </template>
    <template is="dom-if" if="[[stateObj.attributes.wind_speed]]">
      <div class="flex">
        <iron-icon icon="mdi:weather-windy"></iron-icon>
        <div class="main">[[localize('ui.card.weather.attributes.wind_speed')]]</div>
        <div>[[getWind(stateObj.attributes.wind_speed, stateObj.attributes.wind_bearing, localize)]]</div>
      </div>
    </template>
    <template is="dom-if" if="[[stateObj.attributes.visibility]]">
      <div class="flex">
        <iron-icon icon="mdi:eye"></iron-icon>
        <div class="main">[[localize('ui.card.weather.attributes.visibility')]]</div>
        <div>[[stateObj.attributes.visibility]] [[getUnit('length')]]</div>
      </div>
    </template>

    <template is="dom-if" if="[[stateObj.attributes.forecast]]">
      <div class="section">[[localize('ui.card.weather.forecast')]]:</div>
      <template is="dom-repeat" items="[[stateObj.attributes.forecast]]">
        <div class="flex">
          <template is="dom-if" if="[[item.condition]]">
            <iron-icon icon="[[getWeatherIcon(item.condition)]]"></iron-icon>
          </template>
          <div class="main">[[computeDateTime(item.datetime)]]</div>
          <template is="dom-if" if="[[item.templow]]">
            <div class="templow">[[item.templow]] [[getUnit('temperature')]]</div>
          </template>
          <div class="temp">[[item.temperature]] [[getUnit('temperature')]]</div>
        </div>
      </template>
    </template>

    <template is="dom-if" if="stateObj.attributes.attribution">
      <div class="attribution">[[stateObj.attributes.attribution]]</div>
    </template>
`}static get properties(){return{hass:Object,stateObj:Object}}constructor(){super(),this.cardinalDirections=["N","NNE","NE","ENE","E","ESE","SE","SSE","S","SSW","SW","WSW","W","WNW","NW","NNW","N"],this.weatherIcons={"clear-night":"mdi:weather-night",cloudy:"mdi:weather-cloudy",fog:"mdi:weather-fog",hail:"mdi:weather-hail",lightning:"mid:weather-lightning","lightning-rainy":"mdi:weather-lightning-rainy",partlycloudy:"mdi:weather-partlycloudy",pouring:"mdi:weather-pouring",rainy:"mdi:weather-rainy",snowy:"mdi:weather-snowy","snowy-rainy":"mdi:weather-snowy-rainy",sunny:"mdi:weather-sunny",windy:"mdi:weather-windy","windy-variant":"mdi:weather-windy-variant"}}computeDateTime(t){const e=new Date(t),i=this.stateObj.attributes.attribution;return"Powered by Dark Sky"===i||"Data provided by OpenWeatherMap"===i?(new Date).getDay()===e.getDay()?e.toLocaleTimeString(this.hass.selectedLanguage||this.hass.language,{hour:"numeric"}):e.toLocaleDateString(this.hass.selectedLanguage||this.hass.language,{weekday:"long",hour:"numeric"}):e.toLocaleDateString(this.hass.selectedLanguage||this.hass.language,{weekday:"long",month:"short",day:"numeric"})}getUnit(t){return this.hass.config.core.unit_system[t]||""}windBearingToText(t){const e=parseInt(t);return isFinite(e)?this.cardinalDirections[((e+11.25)/22.5|0)%16]:t}getWind(t,e,i){if(null!=e){const a=this.windBearingToText(e);return`${t} ${this.getUnit("length")}/h (${i(`ui.card.weather.cardinal_direction.${a.toLowerCase()}`)||a})`}return`${t} ${this.getUnit("length")}/h`}getWeatherIcon(t){return this.weatherIcons[t]}}),customElements.define("more-info-content",class extends s.a{static get properties(){return{hass:Object,stateObj:Object}}static get observers(){return["stateObjChanged(stateObj, hass)"]}constructor(){super(),this.style.display="block"}stateObjChanged(t,e){let i;t&&e?(this._detachedChild&&(this.appendChild(this._detachedChild),this._detachedChild=null),i=t.attributes&&"custom_ui_more_info"in t.attributes?t.attributes.custom_ui_more_info:"more-info-"+Object(m.a)(t),Object(h.a)(this,i.toUpperCase(),{hass:e,stateObj:t})):this.lastChild&&(this._detachedChild=this.lastChild,this.removeChild(this.lastChild))}});var v=i(39);{const t=["camera","configurator","history_graph"];class e extends(Object(r.a)(s.a)){static get template(){return a["a"]`
    <style include="ha-style-dialog">
      app-toolbar {
        color: var(--more-info-header-color);
        background-color: var(--more-info-header-background);
      }

      app-toolbar [main-title] {
        @apply --ha-more-info-app-toolbar-title;
      }

      state-card-content {
        display: block;
        padding: 16px;
      }

      state-history-charts {
        max-width: 100%;
        margin-bottom: 16px;
      }

      @media all and (min-width: 451px) and (min-height: 501px) {
        .main-title {
          pointer-events: auto;
          cursor: default;
        }
      }

      :host([domain=camera]) paper-dialog-scrollable {
        margin: 0 -24px -5px;
      }
    </style>

    <app-toolbar>
      <paper-icon-button icon="mdi:close" dialog-dismiss=""></paper-icon-button>
      <div class="main-title" main-title="" on-click="enlarge">[[_computeStateName(stateObj)]]</div>
      <template is="dom-if" if="[[canConfigure]]">
        <paper-icon-button icon="mdi:settings" on-click="_gotoSettings"></paper-icon-button>
      </template>
    </app-toolbar>

    <template is="dom-if" if="[[_computeShowStateInfo(stateObj)]]" restamp="">
      <state-card-content state-obj="[[stateObj]]" hass="[[hass]]" in-dialog=""></state-card-content>
    </template>
    <paper-dialog-scrollable dialog-element="[[dialogElement]]">
      <template is="dom-if" if="[[_computeShowHistoryComponent(hass, stateObj)]]" restamp="">
        <ha-state-history-data hass="[[hass]]" filter-type="recent-entity" entity-id="[[stateObj.entity_id]]" data="{{_stateHistory}}" is-loading="{{_stateHistoryLoading}}" cache-config="[[_cacheConfig]]"></ha-state-history-data>
        <state-history-charts history-data="[[_stateHistory]]" is-loading-data="[[_stateHistoryLoading]]" up-to-now="" no-single="[[large]]"></state-history-charts>
      </template>
      <more-info-content state-obj="[[stateObj]]" hass="[[hass]]"></more-info-content>
    </paper-dialog-scrollable>
`}static get properties(){return{hass:Object,stateObj:{type:Object,observer:"_stateObjChanged"},dialogElement:Object,canConfigure:Boolean,domain:{type:String,reflectToAttribute:!0,computed:"_computeDomain(stateObj)"},_stateHistory:Object,_stateHistoryLoading:Boolean,large:{type:Boolean,value:!1,notify:!0},_cacheConfig:{type:Object,value:{refresh:60,cacheKey:null,hoursToShow:24}}}}enlarge(){this.large=!this.large}_computeShowStateInfo(e){return!e||!t.includes(Object(u.a)(e))}_computeShowHistoryComponent(t,e){return t&&e&&Object(b.a)(t,"history")&&!v.c.includes(Object(u.a)(e))}_computeDomain(t){return t?Object(u.a)(t):""}_computeStateName(t){return t?Object(o.a)(t):""}_stateObjChanged(t){t&&this._cacheConfig.cacheKey!==`more_info.${t.entity_id}`&&(this._cacheConfig=Object.assign({},this._cacheConfig,{cacheKey:`more_info.${t.entity_id}`}))}_gotoSettings(){this.fire("more-info-page",{page:"settings"})}}customElements.define("more-info-controls",e)}customElements.define("more-info-settings",class extends(Object(r.a)(s.a)){static get template(){return a["a"]`
    <style>
      app-toolbar {
        color: var(--more-info-header-color);
        background-color: var(--more-info-header-background);

        /* to fit save button */
        padding-right: 0;
      }

      app-toolbar [main-title] {
        @apply --ha-more-info-app-toolbar-title;
      }

      app-toolbar paper-button {
        font-size: .8em;
        margin: 0;
      }

      .form {
        padding: 0 24px 24px;
      }
    </style>

    <app-toolbar>
      <paper-icon-button icon="mdi:arrow-left" on-click="_backTapped"></paper-icon-button>
      <div main-title="">[[_computeStateName(stateObj)]]</div>
      <paper-button on-click="_save">Save</paper-button>
    </app-toolbar>

    <div class="form">
      <paper-input value="{{_name}}" label="Name"></paper-input>
    </div>
`}static get properties(){return{hass:Object,stateObj:Object,_componentLoaded:{type:Boolean,computed:"_computeComponentLoaded(hass)"},registryInfo:{type:Object,observer:"_registryInfoChanged",notify:!0},_name:String}}_computeStateName(t){return t?Object(o.a)(t):""}_computeComponentLoaded(t){return Object(b.a)(t,"config.entity_registry")}_registryInfoChanged(t){this._name=t?t.name:""}_backTapped(){this.fire("more-info-page",{page:null})}_save(){const t={name:this._name};this.hass.callApi("post",`config/entity_registry/${this.stateObj.entity_id}`,t).then(t=>{this.registryInfo=t},()=>{alert("save failed!")})}});var f=i(189);customElements.define("ha-more-info-dialog",class extends(Object(f.a)(s.a)){static get template(){return a["a"]`
    <style include="ha-style-dialog paper-dialog-shared-styles">
      :host {
        font-size: 14px;
        width: 365px;
        border-radius: 2px;
      }

      more-info-controls, more-info-settings {
        --more-info-header-background: var(--secondary-background-color);
        --more-info-header-color: var(--primary-text-color);
        --ha-more-info-app-toolbar-title: {
          /* Design guideline states 24px, changed to 16 to align with state info */
          margin-left: 16px;
          line-height: 1.3em;
          max-height: 2.6em;
          overflow: hidden;
          /* webkit and blink still support simple multiline text-overflow */
          display: -webkit-box;
          -webkit-line-clamp: 2;
          -webkit-box-orient: vertical;
          text-overflow: ellipsis;
        }
      }

      /* overrule the ha-style-dialog max-height on small screens */
      @media all and (max-width: 450px), all and (max-height: 500px) {
        more-info-controls, more-info-settings {
          --more-info-header-background: var(--primary-color);
          --more-info-header-color: var(--text-primary-color);
        }
        :host {
          @apply(--ha-dialog-fullscreen);
        }
        :host::before {
          content: "";
          position: fixed;
          z-index: -1;
          top: 0px;
          left: 0px;
          right: 0px;
          bottom: 0px;
          background-color: inherit;
        }
      }

      :host([data-domain=camera]) {
        width: auto;
      }

      :host([data-domain=history_graph]), :host([large]) {
        width: 90%;
      }
    </style>

    <template is="dom-if" if="[[!_page]]">
      <more-info-controls class="no-padding" hass="[[hass]]" state-obj="[[stateObj]]" dialog-element="[[_dialogElement]]" can-configure="[[_registryInfo]]" large="{{large}}"></more-info-controls>
    </template>
    <template is="dom-if" if="[[_equals(_page, &quot;settings&quot;)]]">
      <more-info-settings class="no-padding" hass="[[hass]]" state-obj="[[stateObj]]" registry-info="{{_registryInfo}}"></more-info-settings>
    </template>
`}static get properties(){return{hass:Object,stateObj:{type:Object,computed:"_computeStateObj(hass)",observer:"_stateObjChanged"},large:{type:Boolean,reflectToAttribute:!0,observer:"_largeChanged"},_dialogElement:Object,_registryInfo:Object,_page:{type:String,value:null},dataDomain:{computed:"_computeDomain(stateObj)",reflectToAttribute:!0}}}static get observers(){return["_dialogOpenChanged(opened)"]}ready(){super.ready(),this._dialogElement=this,this.addEventListener("more-info-page",t=>{this._page=t.detail.page})}_computeDomain(t){return t?Object(u.a)(t):""}_computeStateObj(t){return t.states[t.moreInfoEntityId]||null}_stateObjChanged(t,e){t?(!Object(b.a)(this.hass,"config.entity_registry")||e&&e.entity_id===t.entity_id||this.hass.callApi("get",`config/entity_registry/${t.entity_id}`).then(t=>{this._registryInfo=t},()=>{this._registryInfo=!1}),requestAnimationFrame(()=>requestAnimationFrame(()=>{this.opened=!0}))):this.setProperties({opened:!1,_page:null,_registryInfo:null,large:!1})}_dialogOpenChanged(t){!t&&this.stateObj&&this.fire("hass-more-info",{entityId:null})}_equals(t,e){return t===e}_largeChanged(){this.notifyResize()}})}}]);
//# sourceMappingURL=more-info-dialog-55df7f17d55de7700199.chunk.js.map