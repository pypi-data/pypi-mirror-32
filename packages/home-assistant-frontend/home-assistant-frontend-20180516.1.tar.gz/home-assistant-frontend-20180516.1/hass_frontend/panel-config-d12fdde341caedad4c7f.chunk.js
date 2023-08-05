(window.webpackJsonp=window.webpackJsonp||[]).push([[21],{131:function(e,t,i){"use strict";i(106);var a=i(19);i(2);const s=document.createElement("template");s.setAttribute("style","display: none;"),s.innerHTML='<dom-module id="ha-paper-slider">\n  <template strip-whitespace="">\n    <style include="paper-slider">\n      .pin > .slider-knob > .slider-knob-inner {\n        font-size:  var(--ha-paper-slider-pin-font-size, 10px);\n        line-height: normal;\n      }\n\n      .pin > .slider-knob > .slider-knob-inner::before {\n        top: unset;\n        margin-left: unset;\n\n        bottom: calc(15px + var(--calculated-paper-slider-height)/2);\n        left: 50%;\n        width: 2.6em;\n        height: 2.6em;\n\n        -webkit-transform-origin: left bottom;\n        transform-origin: left bottom;\n        -webkit-transform: rotate(-45deg) scale(0) translate(0);\n        transform: rotate(-45deg) scale(0) translate(0);\n      }\n\n      .pin.expand > .slider-knob > .slider-knob-inner::before {\n        -webkit-transform: rotate(-45deg) scale(1) translate(7px, -7px);\n        transform: rotate(-45deg) scale(1) translate(7px, -7px);\n      }\n\n      .pin > .slider-knob > .slider-knob-inner::after {\n        top: unset;\n        font-size: unset;\n\n        bottom: calc(15px + var(--calculated-paper-slider-height)/2);\n        left: 50%;\n        margin-left: -1.3em;\n        width: 2.6em;\n        height: 2.4em;\n\n        -webkit-transform-origin: center bottom;\n        transform-origin: center bottom;\n        -webkit-transform: scale(0) translate(0);\n        transform: scale(0) translate(0);\n      }\n\n      .pin.expand > .slider-knob > .slider-knob-inner::after {\n        -webkit-transform: scale(1) translate(0, -10px);\n        transform: scale(1) translate(0, -10px);\n      }\n    </style>\n  </template>\n\n\n</dom-module>',document.head.appendChild(s.content);{const e=customElements.get("paper-slider");let t;class i extends e{static get template(){if(!t){t=a.a.import(i.is,"template");const s=document.importNode(e.template.content,!0);s.querySelector("style").remove(),t.content.append(s)}return t}}customElements.define("ha-paper-slider",i)}},139:function(e,t,i){"use strict";i(100),i(97),i(50),i(73),i(83);var a=i(0),s=i(3);i(14),i(131);customElements.define("ha-form",class extends(window.hassMixins.EventsMixin(s.a)){static get template(){return a["a"]`
    <style>
      .error {
        color: red;
      }
    </style>
    <template is="dom-if" if="[[_isArray(schema)]]" restamp="">
      <template is="dom-if" if="[[error.base]]">
          [[computeError(error.base, schema)]]
      </template>

      <template is="dom-repeat" items="[[schema]]">
        <ha-form data="[[_getValue(data, item)]]" schema="[[item]]" error="[[_getValue(error, item)]]" on-data-changed="_valueChanged" compute-label="[[computeLabel]]" compute-error="[[computeError]]"></ha-form>
      </template>
    </template>
    <template is="dom-if" if="[[!_isArray(schema)]]" restamp="">
      <template is="dom-if" if="[[error]]">
        <div class="error">[[computeError(error, schema)]]</div>
      </template>

      <template is="dom-if" if="[[_equals(schema.type, &quot;string&quot;)]]" restamp="">
        <paper-input label="[[computeLabel(schema)]]" value="{{data}}"></paper-input>
      </template>

      <template is="dom-if" if="[[_equals(schema.type, &quot;integer&quot;)]]" restamp="">
        <template is="dom-if" if="[[_isRange(schema)]]" restamp="">
          <div>
            [[computeLabel(schema)]]
            <ha-paper-slider pin="" value="{{data}}" min="[[schema.valueMin]]" max="[[schema.valueMax]]"></ha-paper-slider>
          </div>
        </template>
        <template is="dom-if" if="[[!_isRange(schema)]]" restamp="">
          <paper-input label="[[computeLabel(schema)]]" value="{{data}}" type="number"></paper-input>
        </template>
      </template>

      <template is="dom-if" if="[[_equals(schema.type, &quot;float&quot;)]]" restamp="">
        <!--TODO-->
        <paper-input label="[[computeLabel(schema)]]" value="{{data}}"></paper-input>
      </template>

      <template is="dom-if" if="[[_equals(schema.type, &quot;boolean&quot;)]]" restamp="">
        <paper-checkbox checked="{{data}}">[[computeLabel(schema)]]</paper-checkbox>
      </template>

      <template is="dom-if" if="[[_equals(schema.type, &quot;select&quot;)]]" restamp="">
        <paper-dropdown-menu label="[[computeLabel(schema)]]">
          <paper-listbox slot="dropdown-content" attr-for-selected="item-name" selected="{{data}}">
            <template is="dom-repeat" items="[[schema.options]]">
              <paper-item item-name\$="[[item]]">[[item]]</paper-item>
            </template>
          </paper-listbox>
        </paper-dropdown-menu>
      </template>

    </template>
`}static get properties(){return{data:{type:Object,notify:!0},schema:Object,error:Object,computeLabel:{type:Function,value:()=>e=>e&&e.name},computeError:{type:Function,value:()=>(e,t)=>e}}}_isArray(e){return Array.isArray(e)}_isRange(e){return"valueMin"in e&&"valueMax"in e}_equals(e,t){return e===t}_getValue(e,t){return e[t.name]}_valueChanged(e){this.set(["data",e.model.item.name],e.detail.value)}})},166:function(e,t,i){"use strict";var a=i(0),s=i(3);i(14),i(170);customElements.define("ha-call-service-button",class extends(window.hassMixins.EventsMixin(s.a)){static get template(){return a["a"]`
    <ha-progress-button id="progress" progress="[[progress]]" on-click="buttonTapped"><slot></slot></ha-progress-button>
`}static get properties(){return{hass:{type:Object},progress:{type:Boolean,value:!1},domain:{type:String},service:{type:String},serviceData:{type:Object,value:{}}}}buttonTapped(){this.progress=!0;var e=this,t={domain:this.domain,service:this.service,serviceData:this.serviceData};this.hass.callService(this.domain,this.service,this.serviceData).then(function(){e.progress=!1,e.$.progress.actionSuccess(),t.success=!0},function(){e.progress=!1,e.$.progress.actionError(),t.success=!1}).then(function(){e.fire("hass-service-called",t)})}})},170:function(e,t,i){"use strict";i(59),i(99);var a=i(0),s=i(3);customElements.define("ha-progress-button",class extends s.a{static get template(){return a["a"]`
    <style>
      .container {
        position: relative;
        display: inline-block;
      }

      paper-button {
        transition: all 1s;
      }

      .success paper-button {
        color: white;
        background-color: var(--google-green-500);
        transition: none;
      }

      .error paper-button {
        color: white;
        background-color: var(--google-red-500);
        transition: none;
      }

      paper-button[disabled] {
        color: #c8c8c8;
      }

      .progress {
        @apply --layout;
        @apply --layout-center-center;
        position: absolute;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
      }
    </style>
    <div class="container" id="container">
      <paper-button id="button" disabled="[[computeDisabled(disabled, progress)]]" on-click="buttonTapped">
        <slot></slot>
      </paper-button>
      <template is="dom-if" if="[[progress]]">
        <div class="progress">
          <paper-spinner active=""></paper-spinner>
        </div>
      </template>
    </div>
`}static get properties(){return{hass:{type:Object},progress:{type:Boolean,value:!1},disabled:{type:Boolean,value:!1}}}tempClass(e){var t=this.$.container.classList;t.add(e),setTimeout(()=>{t.remove(e)},1e3)}ready(){super.ready(),this.addEventListener("click",e=>this.buttonTapped(e))}buttonTapped(e){this.progress&&e.stopPropagation()}actionSuccess(){this.tempClass("success")}actionError(){this.tempClass("error")}computeDisabled(e,t){return e||t}})},176:function(e,t,i){"use strict";i(51),i(50),i(134),i(160);var a=i(0),s=i(3),o=(i(177),i(14),i(105),i(15));customElements.define("ha-entity-picker",class extends(window.hassMixins.LocalizeMixin(s.a)){static get template(){return a["a"]`
    <style>
      paper-input > paper-icon-button {
        width: 24px;
        height: 24px;
        padding: 2px;
        color: var(--secondary-text-color);
      }
      [hidden] {
        display: none;
      }
    </style>
    <vaadin-combo-box-light items="[[_states]]" item-value-path="entity_id" item-label-path="entity_id" value="{{value}}" opened="{{opened}}" allow-custom-value="[[allowCustomEntity]]">
      <paper-input autofocus="[[autofocus]]" label="[[_computeLabel(label, localize)]]" class="input" value="[[value]]" disabled="[[disabled]]">
        <paper-icon-button slot="suffix" class="clear-button" icon="mdi:close" no-ripple="" hidden\$="[[!value]]">Clear</paper-icon-button>
        <paper-icon-button slot="suffix" class="toggle-button" icon="[[_computeToggleIcon(opened)]]" hidden="[[!_states.length]]">Toggle</paper-icon-button>
      </paper-input>
      <template>
        <style>
          paper-icon-item {
            margin: -10px;
          }
        </style>
        <paper-icon-item>
          <state-badge state-obj="[[item]]" slot="item-icon"></state-badge>
          <paper-item-body two-line="">
            <div>[[_computeStateName(item)]]</div>
            <div secondary="">[[item.entity_id]]</div>
          </paper-item-body>
        </paper-icon-item>
      </template>
    </vaadin-combo-box-light>
`}static get properties(){return{allowCustomEntity:{type:Boolean,value:!1},hass:{type:Object,observer:"_hassChanged"},_hass:Object,_states:{type:Array,computed:"_computeStates(_hass, domainFilter, entityFilter)"},autofocus:Boolean,label:{type:String},value:{type:String,notify:!0},opened:{type:Boolean,value:!1,observer:"_openedChanged"},domainFilter:{type:String,value:null},entityFilter:{type:Function,value:null},disabled:Boolean}}_computeLabel(e,t){return void 0===e?t("ui.components.entity.entity-picker.entity"):e}_computeStates(e,t,i){if(!e)return[];let a=Object.keys(e.states);t&&(a=a.filter(e=>e.substr(0,e.indexOf("."))===t));let s=a.sort().map(t=>e.states[t]);return i&&(s=s.filter(i)),s}_computeStateName(e){return Object(o.a)(e)}_openedChanged(e){e||(this._hass=this.hass)}_hassChanged(e){this.opened||(this._hass=e)}_computeToggleIcon(e){return e?"mdi:menu-up":"mdi:menu-down"}})},179:function(e,t,i){"use strict";i(51),i(50),i(73);var a=i(0),s=i(3);i(177);customElements.define("ha-combo-box",class extends s.a{static get template(){return a["a"]`
    <style>
      paper-input > paper-icon-button {
        width: 24px;
        height: 24px;
        padding: 2px;
        color: var(--secondary-text-color);
      }
      [hidden] {
        display: none;
      }
    </style>
    <vaadin-combo-box-light items="[[_items]]" item-value-path="[[itemValuePath]]" item-label-path="[[itemLabelPath]]" value="{{value}}" opened="{{opened}}" allow-custom-value="[[allowCustomValue]]">
      <paper-input autofocus="[[autofocus]]" label="[[label]]" class="input" value="[[value]]">
        <paper-icon-button slot="suffix" class="clear-button" icon="mdi:close" hidden\$="[[!value]]">Clear</paper-icon-button>
        <paper-icon-button slot="suffix" class="toggle-button" icon="[[_computeToggleIcon(opened)]]" hidden\$="[[!items.length]]">Toggle</paper-icon-button>
      </paper-input>
      <template>
        <style>
            paper-item {
              margin: -5px -10px;
            }
        </style>
        <paper-item>[[_computeItemLabel(item, itemLabelPath)]]</paper-item>
      </template>
    </vaadin-combo-box-light>
`}static get properties(){return{allowCustomValue:Boolean,items:{type:Object,observer:"_itemsChanged"},_items:Object,itemLabelPath:String,itemValuePath:String,autofocus:Boolean,label:String,opened:{type:Boolean,value:!1,observer:"_openedChanged"},value:{type:String,notify:!0}}}_openedChanged(e){e||(this._items=this.items)}_itemsChanged(e){this.opened||(this._items=e)}_computeToggleIcon(e){return e?"mdi:menu-up":"mdi:menu-down"}_computeItemLabel(e,t){return t?e[t]:e}})},182:function(e,t){window.hassAttributeUtil=window.hassAttributeUtil||{},window.hassAttributeUtil.DOMAIN_DEVICE_CLASS={binary_sensor:["battery","cold","connectivity","door","garage_door","gas","heat","light","lock","moisture","motion","moving","occupancy","opening","plug","power","presence","problem","safety","smoke","sound","vibration","window"],cover:["garage"],sensor:["battery","humidity","illuminance","temperature"]},window.hassAttributeUtil.UNKNOWN_TYPE="json",window.hassAttributeUtil.ADD_TYPE="key-value",window.hassAttributeUtil.TYPE_TO_TAG={string:"ha-customize-string",json:"ha-customize-string",icon:"ha-customize-icon",boolean:"ha-customize-boolean",array:"ha-customize-array","key-value":"ha-customize-key-value"},window.hassAttributeUtil.LOGIC_STATE_ATTRIBUTES=window.hassAttributeUtil.LOGIC_STATE_ATTRIBUTES||{entity_picture:void 0,friendly_name:{type:"string",description:"Name"},icon:{type:"icon"},emulated_hue:{type:"boolean",domains:["emulated_hue"]},emulated_hue_name:{type:"string",domains:["emulated_hue"]},haaska_hidden:void 0,haaska_name:void 0,homebridge_hidden:{type:"boolean"},homebridge_name:{type:"string"},supported_features:void 0,attribution:void 0,custom_ui_more_info:{type:"string"},custom_ui_state_card:{type:"string"},device_class:{type:"array",options:window.hassAttributeUtil.DOMAIN_DEVICE_CLASS,description:"Device class",domains:["binary_sensor","cover","sensor"]},hidden:{type:"boolean",description:"Hide from UI"},assumed_state:{type:"boolean",domains:["switch","light","cover","climate","fan","group"]},initial_state:{type:"string",domains:["automation"]},unit_of_measurement:{type:"string"}}},188:function(e,t,i){"use strict";var a=i(0),s=i(3);i(14),i(179);customElements.define("ha-service-picker",class extends(window.hassMixins.LocalizeMixin(s.a)){static get template(){return a["a"]`
    <ha-combo-box label="[[localize('ui.components.service-picker.service')]]" items="[[_services]]" value="{{value}}" allow-custom-value=""></ha-combo-box>
`}static get properties(){return{hass:{type:Object,observer:"_hassChanged"},_services:Array,value:{type:String,notify:!0}}}_hassChanged(e,t){if(!e)return void(this._services=[]);if(t&&e.config.services===t.config.services)return;const i=[];Object.keys(e.config.services).sort().forEach(t=>{const a=Object.keys(e.config.services[t]).sort();for(let e=0;e<a.length;e++)i.push(`${t}.${a[e]}`)}),this._services=i}})},194:function(e,t){const i=document.createElement("template");i.setAttribute("style","display: none;"),i.innerHTML='<dom-module id="ha-form-style">\n  <template>\n    <style>\n      .form-group {\n        @apply --layout-horizontal;\n        @apply --layout-center;\n        padding: 8px 16px;\n      }\n\n      .form-group label {\n        @apply --layout-flex-2;\n      }\n\n      .form-group .form-control {\n        @apply --layout-flex;\n      }\n\n      .form-group.vertical {\n        @apply --layout-vertical;\n        @apply --layout-start;\n      }\n\n      paper-dropdown-menu.form-control {\n        margin: -9px 0;\n      }\n    </style>\n  </template>\n</dom-module>',document.head.appendChild(i.content)},391:function(e,t,i){"use strict";i.r(t);i(84),i(121);var a=i(0),s=i(3);i(82),i(24),i(59);customElements.define("hass-error-screen",class extends s.a{static get template(){return a["a"]`
    <style include="iron-flex ha-style">
      .placeholder {
        height: 100%;
      }

      .layout {
        height: calc(100% - 64px);
      }

      paper-button {
        font-weight: bold;
        color: var(--primary-color);
      }
    </style>

    <div class="placeholder">
      <app-toolbar>
        <div main-title="">[[title]]</div>
      </app-toolbar>
      <div class="layout vertical center-center">
        <h3>[[error]]</h3>
        <paper-button on-click="backTapped">go back</paper-button>
      </div>
    </div>
`}static get properties(){return{title:{type:String,value:"Home Assistant"},error:{type:String,value:"Oops! It looks like something went wrong."}}}backTapped(){history.back()}});i(14),i(117),i(162),i(129),i(198),i(183),i(51),i(50),i(161),i(160),i(73),i(83),i(104),i(191),i(197);var o=i(157);i(176),i(179),i(133),i(188),i(132),i(130);customElements.define("ha-config-section",class extends s.a{static get template(){return a["a"]`
    <style include="iron-flex ha-style">
      .content {
        padding: 28px 20px 0;
        max-width: 1040px;
        margin: 0 auto;
      }

      .header {
        @apply --paper-font-display1;
        opacity: var(--dark-primary-opacity);
      }

      .together {
        margin-top: 32px;
      }

      .intro {
        @apply --paper-font-subhead;
        width: 100%;
        max-width: 400px;
        margin-right: 40px;
        opacity: var(--dark-primary-opacity);
      }

      .panel {
        margin-top: -24px;
      }

      .panel ::slotted(*) {
        margin-top: 24px;
        display: block;
      }

      .narrow.content {
        max-width: 640px;
      }
      .narrow .together {
        margin-top: 20px;
      }
      .narrow .header {
        @apply --paper-font-headline;
      }
      .narrow .intro {
        font-size: 14px;
        padding-bottom: 20px;
        margin-right: 0;
        max-width: 500px;
      }
    </style>
    <div class\$="[[computeContentClasses(isWide)]]">
      <div class="header"><slot name="header"></slot></div>
      <div class\$="[[computeClasses(isWide)]]">
        <div class="intro">
          <slot name="introduction"></slot>
        </div>
        <div class="panel flex-auto">
          <slot></slot>
        </div>
      </div>
    </div>
`}static get properties(){return{hass:{type:Object},narrow:{type:Boolean},showMenu:{type:Boolean,value:!1},isWide:{type:Boolean,value:!1}}}computeContentClasses(e){return e?"content ":"content narrow"}computeClasses(e){return"together layout "+(e?"horizontal":"vertical narrow")}});class n extends o.a{constructor(e){super(e),this.state.isValid=!0,this.state.value=JSON.stringify(e.value||{},null,2),this.onChange=this.onChange.bind(this)}onChange(e){const t=e.target.value;let i,a;try{i=JSON.parse(t),a=!0}catch(e){a=!1}this.setState({value:t,isValid:a}),a&&this.props.onChange(i)}componentWillReceiveProps({value:e}){e!==this.props.value&&this.setState({value:JSON.stringify(e,null,2),isValid:!0})}render({label:e},{value:t,isValid:i}){const a={minWidth:300,width:"100%"};return i||(a.border="1px solid red"),Object(o.b)("paper-textarea",{label:e,value:t,style:a,"onvalue-changed":this.onChange})}}function r(e,t){const i=this.props[e];if(t.target.value===i[t.target.name])return;const a=Object.assign({},i);t.target.value?a[t.target.name]=t.target.value:delete a[t.target.name],this.props.onChange(this.props.index,a)}class c extends o.a{constructor(){super(),this.onChange=r.bind(this,"trigger"),this.eventDataChanged=this.eventDataChanged.bind(this)}eventDataChanged(e){this.props.onChange(this.props.index,Object.assign({},this.props.trigger,{event_data:e}))}render({trigger:e,localize:t}){const{event_type:i,event_data:a}=e;return Object(o.b)("div",null,Object(o.b)("paper-input",{label:t("ui.panel.config.automation.editor.triggers.type.event.event_type"),name:"event_type",value:i,"onvalue-changed":this.onChange}),Object(o.b)(n,{label:t("ui.panel.config.automation.editor.triggers.type.event.event_data"),value:a,onChange:this.eventDataChanged}))}}c.defaultConfig={event_type:"",event_data:{}};class l extends o.a{constructor(){super(),this.radioGroupPicked=this.radioGroupPicked.bind(this)}radioGroupPicked(e){this.props.onChange(this.props.index,Object.assign({},this.props.trigger,{event:e.target.selected}))}render({trigger:e,localize:t}){const{event:i}=e;return Object(o.b)("div",null,Object(o.b)("label",{id:"eventlabel"},t("ui.panel.config.automation.editor.triggers.type.homeassistant.event")),Object(o.b)("paper-radio-group",{selected:i,"aria-labelledby":"eventlabel","onpaper-radio-group-changed":this.radioGroupPicked},Object(o.b)("paper-radio-button",{name:"start"},t("ui.panel.config.automation.editor.triggers.type.homeassistant.start")),Object(o.b)("paper-radio-button",{name:"shutdown"},t("ui.panel.config.automation.editor.triggers.type.homeassistant.shutdown"))))}}l.defaultConfig={event:"start"};class d extends o.a{constructor(){super(),this.onChange=r.bind(this,"trigger")}render({trigger:e,localize:t}){const{topic:i,payload:a}=e;return Object(o.b)("div",null,Object(o.b)("paper-input",{label:t("ui.panel.config.automation.editor.triggers.type.mqtt.topic"),name:"topic",value:i,"onvalue-changed":this.onChange}),Object(o.b)("paper-input",{label:t("ui.panel.config.automation.editor.triggers.type.mqtt.payload"),name:"payload",value:a,"onvalue-changed":this.onChange}))}}d.defaultConfig={topic:""};class p extends o.a{constructor(){super(),this.onChange=r.bind(this,"trigger"),this.entityPicked=this.entityPicked.bind(this)}entityPicked(e){this.props.onChange(this.props.index,Object.assign({},this.props.trigger,{entity_id:e.target.value}))}render({trigger:e,hass:t,localize:i}){const{value_template:a,entity_id:s,below:n,above:r}=e;return Object(o.b)("div",null,Object(o.b)("ha-entity-picker",{value:s,onChange:this.entityPicked,hass:t,allowCustomEntity:!0}),Object(o.b)("paper-input",{label:i("ui.panel.config.automation.editor.triggers.type.numeric_state.above"),name:"above",value:r,"onvalue-changed":this.onChange}),Object(o.b)("paper-input",{label:i("ui.panel.config.automation.editor.triggers.type.numeric_state.below"),name:"below",value:n,"onvalue-changed":this.onChange}),Object(o.b)("paper-textarea",{label:i("ui.panel.config.automation.editor.triggers.type.numeric_state.value_template"),name:"value_template",value:a,"onvalue-changed":this.onChange}))}}p.defaultConfig={entity_id:""};class u extends o.a{constructor(){super(),this.onChange=r.bind(this,"trigger"),this.entityPicked=this.entityPicked.bind(this)}entityPicked(e){this.props.onChange(this.props.index,Object.assign({},this.props.trigger,{entity_id:e.target.value}))}render({trigger:e,hass:t,localize:i}){const{entity_id:a,to:s}=e,n=e.from,r=e.for;return Object(o.b)("div",null,Object(o.b)("ha-entity-picker",{value:a,onChange:this.entityPicked,hass:t,allowCustomEntity:!0}),Object(o.b)("paper-input",{label:i("ui.panel.config.automation.editor.triggers.type.state.from"),name:"from",value:n,"onvalue-changed":this.onChange}),Object(o.b)("paper-input",{label:i("ui.panel.config.automation.editor.triggers.type.state.to"),name:"to",value:s,"onvalue-changed":this.onChange}),r&&Object(o.b)("pre",null,"For: ",JSON.stringify(r,null,2)))}}u.defaultConfig={entity_id:""};class h extends o.a{constructor(){super(),this.onChange=r.bind(this,"trigger"),this.radioGroupPicked=this.radioGroupPicked.bind(this)}radioGroupPicked(e){this.props.onChange(this.props.index,Object.assign({},this.props.trigger,{event:e.target.selected}))}render({trigger:e,localize:t}){const{offset:i,event:a}=e;return Object(o.b)("div",null,Object(o.b)("label",{id:"eventlabel"},t("ui.panel.config.automation.editor.triggers.type.sun.event")),Object(o.b)("paper-radio-group",{selected:a,"aria-labelledby":"eventlabel","onpaper-radio-group-changed":this.radioGroupPicked},Object(o.b)("paper-radio-button",{name:"sunrise"},t("ui.panel.config.automation.editor.triggers.type.sun.sunrise")),Object(o.b)("paper-radio-button",{name:"sunset"},t("ui.panel.config.automation.editor.triggers.type.sun.sunset"))),Object(o.b)("paper-input",{label:t("ui.panel.config.automation.editor.triggers.type.sun.offset"),name:"offset",value:i,"onvalue-changed":this.onChange}))}}h.defaultConfig={event:"sunrise"};class m extends o.a{constructor(){super(),this.onChange=r.bind(this,"trigger")}render({trigger:e,localize:t}){const{value_template:i}=e;return Object(o.b)("div",null,Object(o.b)("paper-textarea",{label:t("ui.panel.config.automation.editor.triggers.type.template.value_template"),name:"value_template",value:i,"onvalue-changed":this.onChange}))}}m.defaultConfig={value_template:""};class g extends o.a{constructor(){super(),this.onChange=r.bind(this,"trigger")}render({trigger:e,localize:t}){const{at:i}=e;return Object(o.b)("div",null,Object(o.b)("paper-input",{label:t("ui.panel.config.automation.editor.triggers.type.time.at"),name:"at",value:i,"onvalue-changed":this.onChange}))}}function b(e){return"latitude"in e.attributes&&"longitude"in e.attributes}g.defaultConfig={at:""};var f=i(13);function v(e){return b(e)&&"zone"!==Object(f.a)(e)}class y extends o.a{constructor(){super(),this.onChange=r.bind(this,"trigger"),this.radioGroupPicked=this.radioGroupPicked.bind(this),this.entityPicked=this.entityPicked.bind(this),this.zonePicked=this.zonePicked.bind(this)}entityPicked(e){this.props.onChange(this.props.index,Object.assign({},this.props.trigger,{entity_id:e.target.value}))}zonePicked(e){this.props.onChange(this.props.index,Object.assign({},this.props.trigger,{zone:e.target.value}))}radioGroupPicked(e){this.props.onChange(this.props.index,Object.assign({},this.props.trigger,{event:e.target.selected}))}render({trigger:e,hass:t,localize:i}){const{entity_id:a,zone:s,event:n}=e;return Object(o.b)("div",null,Object(o.b)("ha-entity-picker",{label:i("ui.panel.config.automation.editor.triggers.type.zone.entity"),value:a,onChange:this.entityPicked,hass:t,allowCustomEntity:!0,entityFilter:v}),Object(o.b)("ha-entity-picker",{label:i("ui.panel.config.automation.editor.triggers.type.zone.zone"),value:s,onChange:this.zonePicked,hass:t,allowCustomEntity:!0,domainFilter:"zone"}),Object(o.b)("label",{id:"eventlabel"},i("ui.panel.config.automation.editor.triggers.type.zone.event")),Object(o.b)("paper-radio-group",{selected:n,"aria-labelledby":"eventlabel","onpaper-radio-group-changed":this.radioGroupPicked},Object(o.b)("paper-radio-button",{name:"enter"},i("ui.panel.config.automation.editor.triggers.type.zone.enter")),Object(o.b)("paper-radio-button",{name:"leave"},i("ui.panel.config.automation.editor.triggers.type.zone.leave"))))}}y.defaultConfig={entity_id:"",zone:"",event:"enter"};const w={event:c,state:u,homeassistant:l,mqtt:d,numeric_state:p,sun:h,template:m,time:g,zone:y},_=Object.keys(w).sort();class x extends o.a{constructor(){super(),this.typeChanged=this.typeChanged.bind(this)}typeChanged(e){const t=e.target.selectedItem.attributes.platform.value;t!==this.props.trigger.platform&&this.props.onChange(this.props.index,Object.assign({platform:t},w[t].defaultConfig))}render({index:e,trigger:t,onChange:i,hass:a,localize:s}){const n=w[t.platform],r=_.indexOf(t.platform);return n?Object(o.b)("div",null,Object(o.b)("paper-dropdown-menu-light",{label:s("ui.panel.config.automation.editor.triggers.type_select"),"no-animations":!0},Object(o.b)("paper-listbox",{slot:"dropdown-content",selected:r,"oniron-select":this.typeChanged},_.map(e=>Object(o.b)("paper-item",{platform:e},s(`ui.panel.config.automation.editor.triggers.type.${e}.label`))))),Object(o.b)(n,{index:e,trigger:t,onChange:i,hass:a,localize:s})):Object(o.b)("div",null,s("ui.panel.config.automation.editor.triggers.unsupported_platform","platform",t.platform),Object(o.b)("pre",null,JSON.stringify(t,null,2)))}}class C extends o.a{constructor(){super(),this.onDelete=this.onDelete.bind(this)}onDelete(){confirm(this.props.localize("ui.panel.config.automation.editor.triggers.delete_confirm"))&&this.props.onChange(this.props.index,null)}render(e){return Object(o.b)("paper-card",null,Object(o.b)("div",{class:"card-menu"},Object(o.b)("paper-menu-button",{"no-animations":!0,"horizontal-align":"right","horizontal-offset":"-5","vertical-offset":"-5"},Object(o.b)("paper-icon-button",{icon:"mdi:dots-vertical",slot:"dropdown-trigger"}),Object(o.b)("paper-listbox",{slot:"dropdown-content"},Object(o.b)("paper-item",{disabled:!0},e.localize("ui.panel.config.automation.editor.triggers.duplicate")),Object(o.b)("paper-item",{onTap:this.onDelete},e.localize("ui.panel.config.automation.editor.triggers.delete"))))),Object(o.b)("div",{class:"card-content"},Object(o.b)(x,e)))}}class O extends o.a{constructor(){super(),this.addTrigger=this.addTrigger.bind(this),this.triggerChanged=this.triggerChanged.bind(this)}addTrigger(){const e=this.props.trigger.concat(Object.assign({platform:"state"},u.defaultConfig));this.props.onChange(e)}triggerChanged(e,t){const i=this.props.trigger.concat();null===t?i.splice(e,1):i[e]=t,this.props.onChange(i)}render({trigger:e,hass:t,localize:i}){return Object(o.b)("div",{class:"triggers"},e.map((e,a)=>Object(o.b)(C,{index:a,trigger:e,onChange:this.triggerChanged,hass:t,localize:i})),Object(o.b)("paper-card",null,Object(o.b)("div",{class:"card-actions add-card"},Object(o.b)("paper-button",{onTap:this.addTrigger},i("ui.panel.config.automation.editor.triggers.add")))))}}class k extends o.a{constructor(){super(),this.onChange=r.bind(this,"condition"),this.entityPicked=this.entityPicked.bind(this)}entityPicked(e){this.props.onChange(this.props.index,Object.assign({},this.props.condition,{entity_id:e.target.value}))}render({condition:e,hass:t,localize:i}){const{value_template:a,entity_id:s,below:n,above:r}=e;return Object(o.b)("div",null,Object(o.b)("ha-entity-picker",{value:s,onChange:this.entityPicked,hass:t,allowCustomEntity:!0}),Object(o.b)("paper-input",{label:i("ui.panel.config.automation.editor.conditions.type.numeric_state.above"),name:"above",value:r,"onvalue-changed":this.onChange}),Object(o.b)("paper-input",{label:i("ui.panel.config.automation.editor.conditions.type.numeric_state.below"),name:"below",value:n,"onvalue-changed":this.onChange}),Object(o.b)("paper-textarea",{label:i("ui.panel.config.automation.editor.conditions.type.numeric_state.value_template"),name:"value_template",value:a,"onvalue-changed":this.onChange}))}}k.defaultConfig={entity_id:""};class j extends o.a{constructor(){super(),this.onChange=r.bind(this,"condition"),this.entityPicked=this.entityPicked.bind(this)}entityPicked(e){this.props.onChange(this.props.index,Object.assign({},this.props.condition,{entity_id:e.target.value}))}render({condition:e,hass:t,localize:i}){const{entity_id:a,state:s}=e,n=e.for;return Object(o.b)("div",null,Object(o.b)("ha-entity-picker",{value:a,onChange:this.entityPicked,hass:t,allowCustomEntity:!0}),Object(o.b)("paper-input",{label:i("ui.panel.config.automation.editor.conditions.type.state.state"),name:"state",value:s,"onvalue-changed":this.onChange}),n&&Object(o.b)("pre",null,"For: ",JSON.stringify(n,null,2)))}}j.defaultConfig={entity_id:"",state:""};class z extends o.a{constructor(){super(),this.onChange=r.bind(this,"condition"),this.afterPicked=this.radioGroupPicked.bind(this,"after"),this.beforePicked=this.radioGroupPicked.bind(this,"before")}radioGroupPicked(e,t){const i=Object.assign({},this.props.condition);t.target.selected?i[e]=t.target.selected:delete i[e],this.props.onChange(this.props.index,i)}render({condition:e,localize:t}){const{after:i,after_offset:a,before:s,before_offset:n}=e;return Object(o.b)("div",null,Object(o.b)("label",{id:"beforelabel"},t("ui.panel.config.automation.editor.conditions.type.sun.before")),Object(o.b)("paper-radio-group",{"allow-empty-selection":!0,selected:s,"aria-labelledby":"beforelabel","onpaper-radio-group-changed":this.beforePicked},Object(o.b)("paper-radio-button",{name:"sunrise"},t("ui.panel.config.automation.editor.conditions.type.sun.sunrise")),Object(o.b)("paper-radio-button",{name:"sunset"},t("ui.panel.config.automation.editor.conditions.type.sun.sunset"))),Object(o.b)("paper-input",{label:t("ui.panel.config.automation.editor.conditions.type.sun.before_offset"),name:"before_offset",value:n,"onvalue-changed":this.onChange,disabled:void 0===s}),Object(o.b)("label",{id:"afterlabel"},t("ui.panel.config.automation.editor.conditions.type.sun.after")),Object(o.b)("paper-radio-group",{"allow-empty-selection":!0,selected:i,"aria-labelledby":"afterlabel","onpaper-radio-group-changed":this.afterPicked},Object(o.b)("paper-radio-button",{name:"sunrise"},t("ui.panel.config.automation.editor.conditions.type.sun.sunrise")),Object(o.b)("paper-radio-button",{name:"sunset"},t("ui.panel.config.automation.editor.conditions.type.sun.sunset"))),Object(o.b)("paper-input",{label:t("ui.panel.config.automation.editor.conditions.type.sun.after_offset"),name:"after_offset",value:a,"onvalue-changed":this.onChange,disabled:void 0===i}))}}z.defaultConfig={};class A extends o.a{constructor(){super(),this.onChange=r.bind(this,"condition")}render({condition:e,localize:t}){const{value_template:i}=e;return Object(o.b)("div",null,Object(o.b)("paper-textarea",{label:t("ui.panel.config.automation.editor.conditions.type.template.value_template"),name:"value_template",value:i,"onvalue-changed":this.onChange}))}}A.defaultConfig={value_template:""};class S extends o.a{constructor(){super(),this.onChange=r.bind(this,"condition")}render({condition:e,localize:t}){const{after:i,before:a}=e;return Object(o.b)("div",null,Object(o.b)("paper-input",{label:t("ui.panel.config.automation.editor.conditions.type.time.after"),name:"after",value:i,"onvalue-changed":this.onChange}),Object(o.b)("paper-input",{label:t("ui.panel.config.automation.editor.conditions.type.time.before"),name:"before",value:a,"onvalue-changed":this.onChange}))}}function N(e){return b(e)&&"zone"!==Object(f.a)(e)}S.defaultConfig={};class E extends o.a{constructor(){super(),this.onChange=r.bind(this,"condition"),this.entityPicked=this.entityPicked.bind(this),this.zonePicked=this.zonePicked.bind(this)}entityPicked(e){this.props.onChange(this.props.index,Object.assign({},this.props.condition,{entity_id:e.target.value}))}zonePicked(e){this.props.onChange(this.props.index,Object.assign({},this.props.condition,{zone:e.target.value}))}render({condition:e,hass:t,localize:i}){const{entity_id:a,zone:s}=e;return Object(o.b)("div",null,Object(o.b)("ha-entity-picker",{label:i("ui.panel.config.automation.editor.conditions.type.zone.entity"),value:a,onChange:this.entityPicked,hass:t,allowCustomEntity:!0,entityFilter:N}),Object(o.b)("ha-entity-picker",{label:i("ui.panel.config.automation.editor.conditions.type.zone.zone"),value:s,onChange:this.zonePicked,hass:t,allowCustomEntity:!0,domainFilter:"zone"}))}}E.defaultConfig={entity_id:"",zone:""};const P={state:j,numeric_state:k,sun:z,template:A,time:S,zone:E},T=Object.keys(P).sort();class D extends o.a{constructor(){super(),this.typeChanged=this.typeChanged.bind(this)}typeChanged(e){const t=e.target.selectedItem.attributes.condition.value;t!==this.props.condition.condition&&this.props.onChange(this.props.index,Object.assign({condition:t},P[t].defaultConfig))}render({index:e,condition:t,onChange:i,hass:a,localize:s}){const n=P[t.condition],r=T.indexOf(t.condition);return n?Object(o.b)("div",null,Object(o.b)("paper-dropdown-menu-light",{label:s("ui.panel.config.automation.editor.conditions.type_select"),"no-animations":!0},Object(o.b)("paper-listbox",{slot:"dropdown-content",selected:r,"oniron-select":this.typeChanged},T.map(e=>Object(o.b)("paper-item",{condition:e},s(`ui.panel.config.automation.editor.conditions.type.${e}.label`))))),Object(o.b)(n,{index:e,condition:t,onChange:i,hass:a,localize:s})):Object(o.b)("div",null,s("ui.panel.config.automation.editor.conditions.unsupported_condition","condition",t.condition),Object(o.b)("pre",null,JSON.stringify(t,null,2)))}}class I extends o.a{constructor(){super(),this.onDelete=this.onDelete.bind(this)}onDelete(){confirm(this.props.localize("ui.panel.config.automation.editor.conditions.delete_confirm"))&&this.props.onChange(this.props.index,null)}render(e){return Object(o.b)("paper-card",null,Object(o.b)("div",{class:"card-menu"},Object(o.b)("paper-menu-button",{"no-animations":!0,"horizontal-align":"right","horizontal-offset":"-5","vertical-offset":"-5"},Object(o.b)("paper-icon-button",{icon:"mdi:dots-vertical",slot:"dropdown-trigger"}),Object(o.b)("paper-listbox",{slot:"dropdown-content"},Object(o.b)("paper-item",{disabled:!0},e.localize("ui.panel.config.automation.editor.conditions.duplicate")),Object(o.b)("paper-item",{onTap:this.onDelete},e.localize("ui.panel.config.automation.editor.conditions.delete"))))),Object(o.b)("div",{class:"card-content"},Object(o.b)(D,e)))}}class M extends o.a{constructor(){super(),this.addCondition=this.addCondition.bind(this),this.conditionChanged=this.conditionChanged.bind(this)}addCondition(){const e=this.props.condition.concat({condition:"state"});this.props.onChange(e)}conditionChanged(e,t){const i=this.props.condition.concat();null===t?i.splice(e,1):i[e]=t,this.props.onChange(i)}render({condition:e,hass:t,localize:i}){return Object(o.b)("div",{class:"triggers"},e.map((e,a)=>Object(o.b)(I,{index:a,condition:e,onChange:this.conditionChanged,hass:t,localize:i})),Object(o.b)("paper-card",null,Object(o.b)("div",{class:"card-actions add-card"},Object(o.b)("paper-button",{onTap:this.addCondition},i("ui.panel.config.automation.editor.conditions.add")))))}}class L extends o.a{constructor(){super(),this.serviceChanged=this.serviceChanged.bind(this),this.serviceDataChanged=this.serviceDataChanged.bind(this)}serviceChanged(e){this.props.onChange(this.props.index,Object.assign({},this.props.action,{service:e.target.value}))}serviceDataChanged(e){this.props.onChange(this.props.index,Object.assign({},this.props.action,{data:e}))}render({action:e,hass:t,localize:i}){const{service:a,data:s}=e;return Object(o.b)("div",null,Object(o.b)("ha-service-picker",{hass:t,value:a,onChange:this.serviceChanged}),Object(o.b)(n,{label:i("ui.panel.config.automation.editor.actions.type.service.service_data"),value:s,onChange:this.serviceDataChanged}))}}L.defaultConfig={alias:"",service:"",data:{}};class q extends o.a{render({action:e,index:t,onChange:i,hass:a,localize:s}){return Object(o.b)(D,{condition:e,onChange:i,index:t,hass:a,localize:s})}}q.defaultConfig=Object.assign({condition:"state"},j.defaultConfig);class B extends o.a{constructor(){super(),this.onChange=r.bind(this,"action")}render({action:e,localize:t}){const{delay:i}=e;return Object(o.b)("div",null,Object(o.b)("paper-input",{label:t("ui.panel.config.automation.editor.actions.type.delay.delay"),name:"delay",value:i,"onvalue-changed":this.onChange}))}}B.defaultConfig={delay:""};class W extends o.a{constructor(){super(),this.onChange=r.bind(this,"action"),this.serviceDataChanged=this.serviceDataChanged.bind(this)}serviceDataChanged(e){this.props.onChange(this.props.index,Object.assign({},this.props.action,{data:e}))}render({action:e,localize:t}){const{event:i,event_data:a}=e;return Object(o.b)("div",null,Object(o.b)("paper-input",{label:t("ui.panel.config.automation.editor.actions.type.event.event"),name:"event",value:i,"onvalue-changed":this.onChange}),Object(o.b)(n,{label:t("ui.panel.config.automation.editor.actions.type.event.service_data"),value:a,onChange:this.serviceDataChanged}))}}W.defaultConfig={event:"",event_data:{}};class $ extends o.a{constructor(){super(),this.onChange=r.bind(this,"action"),this.onTemplateChange=this.onTemplateChange.bind(this)}onTemplateChange(e){this.props.onChange(this.props.index,Object.assign({},this.props.trigger,{[e.target.name]:e.target.value}))}render({action:e,localize:t}){const{wait_template:i,timeout:a}=e;return Object(o.b)("div",null,Object(o.b)("paper-textarea",{label:t("ui.panel.config.automation.editor.actions.type.wait_template.wait_template"),name:"wait_template",value:i,"onvalue-changed":this.onTemplateChange}),Object(o.b)("paper-input",{label:t("ui.panel.config.automation.editor.actions.type.wait_template.timeout"),name:"timeout",value:a,"onvalue-changed":this.onChange}))}}$.defaultConfig={wait_template:"",timeout:""};const V={service:L,delay:B,wait_template:$,condition:q,event:W},G=Object.keys(V).sort();function U(e){const t=Object.keys(V);for(let i=0;i<t.length;i++)if(t[i]in e)return t[i];return null}class R extends o.a{constructor(){super(),this.typeChanged=this.typeChanged.bind(this)}typeChanged(e){const t=e.target.selectedItem.attributes.action.value;U(this.props.action)!==t&&this.props.onChange(this.props.index,V[t].defaultConfig)}render({index:e,action:t,onChange:i,hass:a,localize:s}){const n=U(t),r=n&&V[n],c=G.indexOf(n);return r?Object(o.b)("div",null,Object(o.b)("paper-dropdown-menu-light",{label:s("ui.panel.config.automation.editor.actions.type_select"),"no-animations":!0},Object(o.b)("paper-listbox",{slot:"dropdown-content",selected:c,"oniron-select":this.typeChanged},G.map(e=>Object(o.b)("paper-item",{action:e},s(`ui.panel.config.automation.editor.actions.type.${e}.label`))))),Object(o.b)(r,{index:e,action:t,onChange:i,hass:a,localize:s})):Object(o.b)("div",null,s("ui.panel.config.automation.editor.actions.unsupported_action","action",n),Object(o.b)("pre",null,JSON.stringify(t,null,2)))}}class H extends o.a{constructor(){super(),this.onDelete=this.onDelete.bind(this)}onDelete(){confirm(this.props.localize("ui.panel.config.automation.editor.actions.delete_confirm"))&&this.props.onChange(this.props.index,null)}render(e){return Object(o.b)("paper-card",null,Object(o.b)("div",{class:"card-menu"},Object(o.b)("paper-menu-button",{"no-animations":!0,"horizontal-align":"right","horizontal-offset":"-5","vertical-offset":"-5"},Object(o.b)("paper-icon-button",{icon:"mdi:dots-vertical",slot:"dropdown-trigger"}),Object(o.b)("paper-listbox",{slot:"dropdown-content"},Object(o.b)("paper-item",{disabled:!0},e.localize("ui.panel.config.automation.editor.actions.duplicate")),Object(o.b)("paper-item",{onTap:this.onDelete},e.localize("ui.panel.config.automation.editor.actions.delete"))))),Object(o.b)("div",{class:"card-content"},Object(o.b)(R,e)))}}class F extends o.a{constructor(){super(),this.addAction=this.addAction.bind(this),this.actionChanged=this.actionChanged.bind(this)}addAction(){const e=this.props.script.concat({service:""});this.props.onChange(e)}actionChanged(e,t){const i=this.props.script.concat();null===t?i.splice(e,1):i[e]=t,this.props.onChange(i)}render({script:e,hass:t,localize:i}){return Object(o.b)("div",{class:"script"},e.map((e,a)=>Object(o.b)(H,{index:a,action:e,onChange:this.actionChanged,hass:t,localize:i})),Object(o.b)("paper-card",null,Object(o.b)("div",{class:"card-actions add-card"},Object(o.b)("paper-button",{onTap:this.addAction},i("ui.panel.config.automation.editor.actions.add")))))}}class J extends o.a{constructor(){super(),this.onChange=this.onChange.bind(this),this.triggerChanged=this.triggerChanged.bind(this),this.conditionChanged=this.conditionChanged.bind(this),this.actionChanged=this.actionChanged.bind(this)}onChange(e){this.props.onChange(Object.assign({},this.props.automation,{[e.target.name]:e.target.value}))}triggerChanged(e){this.props.onChange(Object.assign({},this.props.automation,{trigger:e}))}conditionChanged(e){this.props.onChange(Object.assign({},this.props.automation,{condition:e}))}actionChanged(e){this.props.onChange(Object.assign({},this.props.automation,{action:e}))}render({automation:e,isWide:t,hass:i,localize:a}){const{alias:s,trigger:n,condition:r,action:c}=e;return Object(o.b)("div",null,Object(o.b)("ha-config-section",{"is-wide":t},Object(o.b)("span",{slot:"header"},s),Object(o.b)("span",{slot:"introduction"},a("ui.panel.config.automation.editor.introduction")),Object(o.b)("paper-card",null,Object(o.b)("div",{class:"card-content"},Object(o.b)("paper-input",{label:a("ui.panel.config.automation.editor.alias"),name:"alias",value:s,"onvalue-changed":this.onChange})))),Object(o.b)("ha-config-section",{"is-wide":t},Object(o.b)("span",{slot:"header"},a("ui.panel.config.automation.editor.triggers.header")),Object(o.b)("span",{slot:"introduction"},Object(o.b)("ha-markdown",{content:a("ui.panel.config.automation.editor.triggers.introduction")})),Object(o.b)(O,{trigger:n,onChange:this.triggerChanged,hass:i,localize:a})),Object(o.b)("ha-config-section",{"is-wide":t},Object(o.b)("span",{slot:"header"},a("ui.panel.config.automation.editor.conditions.header")),Object(o.b)("span",{slot:"introduction"},Object(o.b)("ha-markdown",{content:a("ui.panel.config.automation.editor.conditions.introduction")})),Object(o.b)(M,{condition:r||[],onChange:this.conditionChanged,hass:i,localize:a})),Object(o.b)("ha-config-section",{"is-wide":t},Object(o.b)("span",{slot:"header"},a("ui.panel.config.automation.editor.actions.header")),Object(o.b)("span",{slot:"introduction"},Object(o.b)("ha-markdown",{content:a("ui.panel.config.automation.editor.actions.introduction")})),Object(o.b)(F,{script:c,onChange:this.actionChanged,hass:i,localize:a})))}}function Y(e){Object(o.c)(()=>null,e)}var Z=i(15);customElements.define("ha-automation-editor",class extends(window.hassMixins.LocalizeMixin(window.hassMixins.NavigateMixin(s.a))){static get template(){return a["a"]`
    <style include="ha-style">
      .errors {
        padding: 20px;
        font-weight: bold;
        color: var(--google-red-500);
      }
      .content {
        padding-bottom: 20px;
      }
      paper-card {
        display: block;
      }
      .triggers,
      .script {
        margin-top: -16px;
      }
      .triggers paper-card,
      .script paper-card {
        margin-top: 16px;
      }
      .add-card paper-button {
        display: block;
        text-align: center;
      }
      .card-menu {
        position: absolute;
        top: 0;
        right: 0;
        z-index: 1;
        color: var(--primary-text-color);
      }
      .card-menu paper-item {
        cursor: pointer;
      }
      span[slot=introduction] a {
        color: var(--primary-color);
      }
      paper-fab {
        position: fixed;
        bottom: 16px;
        right: 16px;
        z-index: 1;
        margin-bottom: -80px;
        transition: margin-bottom .3s;
      }

      paper-fab[is-wide] {
        bottom: 24px;
        right: 24px;
      }

      paper-fab[dirty] {
        margin-bottom: 0;
      }
    </style>

    <ha-app-layout has-scrolling-region="">
      <app-header slot="header" fixed="">
        <app-toolbar>
          <paper-icon-button icon="mdi:arrow-left" on-click="backTapped"></paper-icon-button>
          <div main-title="">[[name]]</div>
        </app-toolbar>
      </app-header>

      <div class="content">
        <template is="dom-if" if="[[errors]]">
          <div class="errors">[[errors]]</div>
        </template>
        <div id="root"></div>
      </div>
      <paper-fab slot="fab" is-wide\$="[[isWide]]" dirty\$="[[dirty]]" icon="mdi:content-save" title="[[localize('ui.panel.config.automation.editor.save')]]" on-click="saveAutomation"></paper-fab>
    </ha-app-layout>
`}static get properties(){return{hass:{type:Object,observer:"_updateComponent"},narrow:{type:Boolean},showMenu:{type:Boolean,value:!1},errors:{type:Object,value:null},dirty:{type:Boolean,value:!1},config:{type:Object,value:null},automation:{type:Object,observer:"automationChanged"},creatingNew:{type:Boolean,observer:"creatingNewChanged"},name:{type:String,computed:"computeName(automation, localize)"},isWide:{type:Boolean,observer:"_updateComponent"},_rendered:{type:Object,value:null},_renderScheduled:{type:Boolean,value:!1}}}ready(){this.configChanged=this.configChanged.bind(this),super.ready()}disconnectedCallback(){super.disconnectedCallback(),this._rendered&&(Y(this._rendered),this._rendered=null)}configChanged(e){null!==this._rendered&&(this.config=e,this.errors=null,this.dirty=!0,this._updateComponent())}automationChanged(e,t){e&&(this.hass?t&&t.attributes.id===e.attributes.id||this.hass.callApi("get","config/automation/config/"+e.attributes.id).then(function(e){["trigger","condition","action"].forEach(function(t){var i=e[t];i&&!Array.isArray(i)&&(e[t]=[i])}),this.dirty=!1,this.config=e,this._updateComponent()}.bind(this)):setTimeout(()=>this.automationChanged(e,t),0))}creatingNewChanged(e){e&&(this.dirty=!1,this.config={alias:this.localize("ui.panel.config.automation.editor.default_name"),trigger:[{platform:"state"}],condition:[],action:[{service:""}]},this._updateComponent())}backTapped(){this.dirty&&!confirm(this.localize("ui.panel.config.automation.editor.unsaved_confirm"))||history.back()}_updateComponent(){var e,t,i;!this._renderScheduled&&this.hass&&this.config&&(this._renderScheduled=!0,Promise.resolve().then(()=>{this._rendered=(e=this.$.root,t={automation:this.config,onChange:this.configChanged,isWide:this.isWide,hass:this.hass,localize:this.localize},i=this._rendered,Object(o.c)(Object(o.b)(J,t),e,i)),this._renderScheduled=!1}))}saveAutomation(){var e=this.creatingNew?""+Date.now():this.automation.attributes.id;this.hass.callApi("post","config/automation/config/"+e,this.config).then(function(){this.dirty=!1,this.creatingNew&&this.navigate(`/config/automation/edit/${e}`,!0)}.bind(this),function(e){throw this.errors=e.body.message,e}.bind(this))}computeName(e,t){return e?Object(Z.a)(e):t("ui.panel.config.automation.editor.default_name")}});customElements.define("ha-automation-picker",class extends(window.hassMixins.LocalizeMixin(window.hassMixins.NavigateMixin(s.a))){static get template(){return a["a"]`
    <style include="ha-style">
      :host {
        display: block;
      }

      paper-item {
        cursor: pointer;
      }

      paper-fab {
        position: fixed;
        bottom: 16px;
        right: 16px;
        z-index: 1;
      }

      paper-fab[is-wide] {
        bottom: 24px;
        right: 24px;
      }

      a {
        color: var(--primary-color);
      }

      ha-markdown p {
        margin: 0px;
      }
    </style>

    <ha-app-layout has-scrolling-region="">
      <app-header slot="header" fixed="">
        <app-toolbar>
          <paper-icon-button icon="mdi:arrow-left" on-click="_backTapped"></paper-icon-button>
          <div main-title="">[[localize('ui.panel.config.automation.caption')]]</div>
        </app-toolbar>
      </app-header>

      <ha-config-section is-wide="[[isWide]]">
        <div slot="header">[[localize('ui.panel.config.automation.picker.header')]]</div>
        <div slot="introduction">
          <ha-markdown content="[[localize('ui.panel.config.automation.picker.introduction')]]"></ha-markdown>
        </div>

        <paper-card heading="[[localize('ui.panel.config.automation.picker.pick_automation')]]">
          <template is="dom-if" if="[[!automations.length]]">
            <div class="card-content">
              <p>[[localize('ui.panel.config.automation.picker.no_automations')]]</p>
            </div>
          </template>
          <template is="dom-repeat" items="[[automations]]" as="automation">
            <paper-item>
              <paper-item-body two-line="" on-click="automationTapped">
                <div>[[computeName(automation)]]</div>
                <div secondary="">[[computeDescription(automation)]]</div>
              </paper-item-body>
              <iron-icon icon="mdi:chevron-right"></iron-icon>
            </paper-item>
          </template>
        </paper-card>
      </ha-config-section>

      <paper-fab slot="fab" is-wide\$="[[isWide]]" icon="mdi:plus" title="[[localize('ui.panel.config.automation.picker.add_automation')]]" on-click="addAutomation"></paper-fab>
    </ha-app-layout>
`}static get properties(){return{hass:{type:Object},narrow:{type:Boolean},showMenu:{type:Boolean,value:!1},automations:{type:Array},isWide:{type:Boolean}}}automationTapped(e){this.navigate("/config/automation/edit/"+this.automations[e.model.index].attributes.id)}addAutomation(){this.navigate("/config/automation/new")}computeName(e){return Object(Z.a)(e)}computeDescription(e){return""}_backTapped(){history.back()}});customElements.define("ha-config-automation",class extends s.a{static get template(){return a["a"]`
    <style>
      ha-automation-picker,
      ha-automation-editor {
        height: 100%;
      }
    </style>
    <app-route route="[[route]]" pattern="/automation/edit/:automation" data="{{_routeData}}" active="{{_edittingAutomation}}"></app-route>
    <app-route route="[[route]]" pattern="/automation/new" active="{{_creatingNew}}"></app-route>

    <template is="dom-if" if="[[!showEditor]]">
      <ha-automation-picker hass="[[hass]]" narrow="[[narrow]]" show-menu="[[showMenu]]" automations="[[automations]]" is-wide="[[isWide]]"></ha-automation-picker>
    </template>

    <template is="dom-if" if="[[showEditor]]" restamp="">
      <ha-automation-editor hass="[[hass]]" automation="[[automation]]" is-wide="[[isWide]]" creating-new="[[_creatingNew]]"></ha-automation-editor>
    </template>
`}static get properties(){return{hass:Object,narrow:Boolean,showMenu:Boolean,route:Object,isWide:Boolean,_routeData:Object,_routeMatches:Boolean,_creatingNew:Boolean,_edittingAutomation:Boolean,automations:{type:Array,computed:"computeAutomations(hass)"},automation:{type:Object,computed:"computeAutomation(automations, _edittingAutomation, _routeData)"},showEditor:{type:Boolean,computed:"computeShowEditor(_edittingAutomation, _creatingNew)"}}}computeAutomation(e,t,i){if(!e||!t)return null;for(var a=0;a<e.length;a++)if(e[a].attributes.id===i.automation)return e[a];return null}computeAutomations(e){var t=[];return Object.keys(e.states).forEach(function(i){var a=e.states[i];"automation"===Object(f.a)(a)&&"id"in a.attributes&&t.push(a)}),t.sort(function(e,t){var i=(e.attributes.alias||e.entity_id).toLowerCase(),a=(t.attributes.alias||t.entity_id).toLowerCase();return i<a?-1:i>a?1:0})}computeShowEditor(e,t){return t||e}});var K=i(7),Q=i(16);i(170);customElements.define("ha-call-api-button",class extends(window.hassMixins.EventsMixin(s.a)){static get template(){return a["a"]`
    <ha-progress-button id="progress" progress="[[progress]]" on-click="buttonTapped" disabled="[[disabled]]"><slot></slot></ha-progress-button>
`}static get properties(){return{hass:Object,progress:{type:Boolean,value:!1},path:String,method:{type:String,value:"POST"},data:{type:Object,value:{}},disabled:{type:Boolean,value:!1}}}buttonTapped(){this.progress=!0;const e={method:this.method,path:this.path,data:this.data};this.hass.callApi(this.method,this.path,this.data).then(t=>{this.progress=!1,this.$.progress.actionSuccess(),e.success=!0,e.response=t},t=>{this.progress=!1,this.$.progress.actionError(),e.success=!1,e.response=t}).then(()=>{this.fire("hass-api-called",e)})}});i(119);customElements.define("hass-subpage",class extends s.a{static get template(){return a["a"]`
    <style include="ha-style"></style>
    <app-header-layout has-scrolling-region="">
      <app-header slot="header" fixed="">
        <app-toolbar>
          <paper-icon-button icon="mdi:arrow-left" on-click="_backTapped"></paper-icon-button>
          <div main-title="">[[header]]</div>
        </app-toolbar>
      </app-header>

      <slot></slot>
    </app-header-layout>
`}static get properties(){return{header:String}}_backTapped(){history.back()}});var X=i(62);customElements.define("ha-config-cloud-account",class extends(window.hassMixins.EventsMixin(s.a)){static get template(){return a["a"]`
    <style include="iron-flex ha-style">
      .content {
        padding-bottom: 24px;
      }
      paper-card {
        display: block;
      }
      .account-row {
        display: flex;
        padding: 0 16px;
      }
      paper-button {
        align-self: center;
      }
      .soon {
        font-style: italic;
        margin-top: 24px;
        text-align: center;
      }
      .nowrap {
        white-space: nowrap;;
      }
      .wrap {
        white-space: normal;
      }
      .status {
        text-transform: capitalize;
        padding: 16px;
      }
      paper-button {
        color: var(--primary-color);
        font-weight: 500;
      }
      a {
        color: var(--primary-color);
      }
    </style>
    <hass-subpage header="Cloud Account">
      <div class="content">
        <ha-config-section is-wide="[[isWide]]">
          <span slot="header">Home Assistant Cloud</span>
          <span slot="introduction">
            Thank you for supporting Home Assistant. It's because of people like you that we are able to run this project and make a great home automation experience for everyone. Thank you!
          </span>

          <paper-card heading="Account">
            <div class="account-row">
              <paper-item-body two-line="">
                [[account.email]]
                <div secondary="" class="wrap">
                  <span class="nowrap">Subscription expires on </span>
                  <span class="nowrap">[[_formatExpiration(account.sub_exp)]]</span>
                </div>
              </paper-item-body>
              <paper-button on-click="handleLogout">Sign out</paper-button>
            </div>

            <div class="account-row">
              <paper-item-body>
                Cloud connection status
              </paper-item-body>
              <div class="status">[[account.cloud]]</div>
            </div>
          </paper-card>
        </ha-config-section>

        <ha-config-section is-wide="[[isWide]]">
          <span slot="header">Integrations</span>
          <span slot="introduction">
            Integrations for Home Assistant Cloud allow you to connect with services in the cloud
            without having to expose your Home Assistant instance publicly on the internet.
          </span>

          <paper-card heading="Alexa">
            <div class="card-content">
              With the Alexa integration for Home Assistant Cloud you'll be able to control all your Home Assistant devices via any Alexa-enabled device.
              <ul>
                <li>
                  <a href="https://alexa.amazon.com/spa/index.html#skills/dp/B0772J1QKB/?ref=skill_dsk_skb_sr_2" target="_blank">
                    Activate the Home Assistant skill for Alexa
                  </a>
                </li>
                <li>
                  <a href="https://www.home-assistant.io/cloud/alexa/" target="_blank">
                    Config documentation
                  </a>
                </li>
              </ul>
              <p><em>This integration requires an Alexa-enabled device like the Amazon Echo.</em></p>
            </div>
          </paper-card>

          <paper-card heading="Google Assistant">
            <div class="card-content">
              With the Google Assistant integration for Home Assistant Cloud you'll be able to control all your Home Assistant devices via any Google Assistant-enabled device.
              <ul>
                <li>
                  <a href="https://assistant.google.com/services/a/uid/00000091fd5fb875" target="_blank">
                    Activate the Home Assistant skill for Google Assistant
                  </a>
                </li>
                <li>
                  <a href="https://www.home-assistant.io/cloud/google_assistant/" target="_blank">
                    Config documentation
                  </a>
                </li>
              </ul>
              <p><em>This integration requires a Google Assistant-enabled device like the Google Home or Android phone.</em></p>
            </div>
            <div class="card-actions">
              <ha-call-api-button hass="[[hass]]" path="cloud/google_actions/sync">Sync devices</ha-call-api-button>
            </div>
          </paper-card>
        </ha-config-section>
      </div>
    </hass-subpage>
`}static get properties(){return{hass:Object,isWide:Boolean,account:{type:Object,observer:"_accountChanged"}}}handleLogout(){this.hass.callApi("post","cloud/logout").then(()=>this.fire("ha-account-refreshed",{account:null}))}_formatExpiration(e){return Object(X.a)(new Date(e))}_accountChanged(e){e&&"connecting"===e.cloud?this._accountUpdater||setTimeout(()=>{this._accountUpdater=null,this.hass.callApi("get","cloud/account").then(e=>this.fire("ha-account-refreshed",{account:e}))},5e3):this._accountUpdater&&(clearTimeout(this._accountUpdater),this._accountUpdater=null)}});customElements.define("ha-config-cloud-forgot-password",class extends(window.hassMixins.EventsMixin(s.a)){static get template(){return a["a"]`
    <style include="iron-flex ha-style">
      .content {
        padding-bottom: 24px;
      }

      paper-card {
        display: block;
        max-width: 600px;
        margin: 0 auto;
        margin-top: 24px;
      }
      h1 {
        @apply --paper-font-headline;
        margin: 0;
      }
      .error {
        color: var(--google-red-500);
      }
      .card-actions {
        display: flex;
        justify-content: space-between;
        align-items: center;
      }
      .card-actions a {
        color: var(--primary-text-color);
      }
      [hidden] {
        display: none;
      }
    </style>
    <hass-subpage header="Forgot Password">
      <div class="content">
        <paper-card>
          <div class="card-content">
            <h1>Forgot your password?</h1>
            <p>
              Enter your email address and we will send you a link to reset your password.
            </p>
            <div class="error" hidden\$="[[!_error]]">[[_error]]</div>
            <paper-input autofocus="" id="email" label="E-mail" value="{{email}}" type="email" on-keydown="_keyDown" error-message="Invalid email"></paper-input>
          </div>
          <div class="card-actions">
            <ha-progress-button on-click="_handleEmailPasswordReset" progress="[[_requestInProgress]]">Send reset email</ha-progress-button>
          </div>
        </paper-card>
      </div>
    </hass-subpage>
`}static get properties(){return{hass:Object,email:{type:String,notify:!0,observer:"_emailChanged"},_requestInProgress:{type:Boolean,value:!1},_error:{type:String,value:""}}}_emailChanged(){this._error="",this.$.email.invalid=!1}_keyDown(e){13===e.keyCode&&(this._handleEmailPasswordReset(),e.preventDefault())}_handleEmailPasswordReset(){this.email&&this.email.includes("@")||(this.$.email.invalid=!0),this.$.email.invalid||(this._requestInProgress=!0,this.hass.callApi("post","cloud/forgot_password",{email:this.email}).then(()=>{this._requestInProgress=!1,this.fire("cloud-done",{flashMessage:"Check your email for instructions on how to reset your password."})},e=>this.setProperties({_requestInProgress:!1,_error:e&&e.body&&e.body.message?e.body.message:"Unknown error"})))}});i(77);customElements.define("ha-config-cloud-login",class extends(window.hassMixins.NavigateMixin(window.hassMixins.EventsMixin(s.a))){static get template(){return a["a"]`
    <style include="iron-flex ha-style">
      .content {
        padding-bottom: 24px;
      }
      [slot=introduction] a {
        color: var(--primary-color);
      }
      paper-card {
        display: block;
      }
      paper-item {
        cursor: pointer;
      }
      paper-card:last-child {
        margin-top: 24px;
      }
      h1 {
        @apply --paper-font-headline;
        margin: 0;
      }
      .error {
        color: var(--google-red-500);
      }
      .card-actions {
        display: flex;
        justify-content: space-between;
        align-items: center;
      }
      [hidden] {
        display: none;
      }
      .flash-msg {
        padding-right: 44px;
      }
      .flash-msg paper-icon-button {
        position: absolute;
        top: 8px;
        right: 8px;
        color: var(--secondary-text-color);
      }
    </style>
    <hass-subpage header="Cloud Login">
      <div class="content">
        <ha-config-section is-wide="[[isWide]]">
          <span slot="header">Home Assistant Cloud</span>
          <span slot="introduction">
            The Home Assistant Cloud allows your local Home Assistant instance to connect with cloud-only services like Amazon Alexa.
            <p><a href="https://www.home-assistant.io/components/cloud/" target="_blank">Learn more</a></p>
          </span>

          <paper-card hidden\$="[[!flashMessage]]">
            <div class="card-content flash-msg">
              [[flashMessage]]
              <paper-icon-button icon="mdi:close" on-click="_dismissFlash">Dismiss</paper-icon-button>
              <paper-ripple id="flashRipple" noink=""></paper-ripple>
            </div>
          </paper-card>

          <paper-card>
            <div class="card-content">
              <h1>Sign In</h1>
              <div class="error" hidden\$="[[!_error]]">[[_error]]</div>
              <paper-input label="Email" id="email" type="email" value="{{email}}" on-keydown="_keyDown" error-message="Invalid email"></paper-input>
              <paper-input id="password" label="Password" value="{{_password}}" type="password" on-keydown="_keyDown" error-message="Passwords are at least 8 characters"></paper-input>
            </div>
            <div class="card-actions">
              <ha-progress-button on-click="_handleLogin" progress="[[_requestInProgress]]">Sign in</ha-progress-button>
              <button class="link" hidden="[[_requestInProgress]]" on-click="_handleForgotPassword">forgot password?</button>
            </div>
          </paper-card>

          <paper-card>
            <paper-item on-click="_handleRegister">
              <paper-item-body two-line="">
                Create Account
                <div secondary="">Get up and running quickly.</div>
              </paper-item-body>
              <iron-icon icon="mdi:chevron-right"></iron-icon>
            </paper-item>
          </paper-card>
        </ha-config-section>
      </div>
    </hass-subpage>
`}static get properties(){return{hass:Object,isWide:Boolean,email:{type:String,notify:!0},_password:{type:String,value:""},_requestInProgress:{type:Boolean,value:!1},flashMessage:{type:String,notify:!0},_error:String}}static get observers(){return["_inputChanged(email, _password)"]}connectedCallback(){super.connectedCallback(),this.flashMessage&&requestAnimationFrame(()=>requestAnimationFrame(()=>this.$.flashRipple.simulatedRipple()))}_inputChanged(){this.$.email.invalid=!1,this.$.password.invalid=!1,this._error=!1}_keyDown(e){13===e.keyCode&&(this._handleLogin(),e.preventDefault())}_handleLogin(){let e=!1;this.email&&this.email.includes("@")||(this.$.email.invalid=!0,this.$.email.focus(),e=!0),this._password.length<8&&(this.$.password.invalid=!0,e||(e=!0,this.$.password.focus())),e||(this._requestInProgress=!0,this.hass.callApi("post","cloud/login",{email:this.email,password:this._password}).then(e=>{this.fire("ha-account-refreshed",{account:e}),this.setProperties({email:"",_password:""})},e=>{this._password="";const t=e&&e.body&&e.body.code;if("PasswordChangeRequired"===t)return alert("You need to change your password before logging in."),void this.navigate("/config/cloud/forgot-password");const i={_requestInProgress:!1,_error:e&&e.body&&e.body.message?e.body.message:"Unknown error"};"UserNotConfirmed"===t&&(i._error="You need to confirm your email before logging in."),this.setProperties(i),this.$.email.focus()}))}_handleRegister(){this.flashMessage="",this.navigate("/config/cloud/register")}_handleForgotPassword(){this.flashMessage="",this.navigate("/config/cloud/forgot-password")}_dismissFlash(){setTimeout(()=>{this.flashMessage=""},200)}});customElements.define("ha-config-cloud-register",class extends(window.hassMixins.EventsMixin(s.a)){static get template(){return a["a"]`
    <style include="iron-flex ha-style">
      a {
        color: var(--primary-color);
      }
      paper-card {
        display: block;
      }
      paper-item {
        cursor: pointer;
      }
      paper-card:last-child {
        margin-top: 24px;
      }
      h1 {
        @apply --paper-font-headline;
        margin: 0;
      }
      .error {
        color: var(--google-red-500);
      }
      .card-actions {
        display: flex;
        justify-content: space-between;
        align-items: center;
      }
      [hidden] {
        display: none;
      }
    </style>
    <hass-subpage header="Register Account">
      <div class="content">
        <ha-config-section is-wide="[[isWide]]">
          <span slot="header">Register with the Home Assistant Cloud</span>
          <span slot="introduction">
            Register today to easily connect with the Home Assistant Cloud. This will allow you to unlock great new services and functionality, like Amazon Alexa integration.

            <p>
              By registering an account you agree to the following terms and conditions.
              </p><ul>
                <li><a href="https://home-assistant.io/tos/" target="_blank">Terms and Conditions</a></li>
                <li><a href="https://home-assistant.io/privacy/" target="_blank">Privacy Policy</a></li>
              </ul>
            <p></p>
          </span>

          <paper-card>
            <div class="card-content">
              <div class="header">
                <h1>Register</h1>
                <div class="error" hidden\$="[[!_error]]">[[_error]]</div>
              </div>
              <paper-input autofocus="" id="email" label="Email address" type="email" value="{{email}}" on-keydown="_keyDown" error-message="Invalid email"></paper-input>
              <paper-input id="password" label="Password" value="{{_password}}" type="password" on-keydown="_keyDown" error-message="Your password needs to be at least 8 characters"></paper-input>
            </div>
            <div class="card-actions">
              <ha-progress-button on-click="_handleRegister" progress="[[_requestInProgress]]">Create Account</ha-progress-button>
              <button class="link" hidden="[[_requestInProgress]]" on-click="_handleResendVerifyEmail">Resend confirmation email</button>
            </div>
          </paper-card>
        </ha-config-section>
      </div>
    </hass-subpage>
`}static get properties(){return{hass:Object,isWide:Boolean,email:{type:String,notify:!0},_requestInProgress:{type:Boolean,value:!1},_password:{type:String,value:""},_error:{type:String,value:""}}}static get observers(){return["_inputChanged(email, _password)"]}_inputChanged(){this._error="",this.$.email.invalid=!1,this.$.password.invalid=!1}_keyDown(e){13===e.keyCode&&(this._handleRegister(),e.preventDefault())}_handleRegister(){let e=!1;this.email&&this.email.includes("@")||(this.$.email.invalid=!0,this.$.email.focus(),e=!0),this._password.length<8&&(this.$.password.invalid=!0,e||(e=!0,this.$.password.focus())),e||(this._requestInProgress=!0,this.hass.callApi("post","cloud/register",{email:this.email,password:this._password}).then(()=>this._verificationEmailSent(),e=>{this._password="",this.setProperties({_requestInProgress:!1,_error:e&&e.body&&e.body.message?e.body.message:"Unknown error"})}))}_handleResendVerifyEmail(){this.email?this.hass.callApi("post","cloud/resend_confirm",{email:this.email}).then(()=>this._verificationEmailSent(),e=>this.setProperties({_error:e&&e.body&&e.body.message?e.body.message:"Unknown error"})):this.$.email.invalid=!0}_verificationEmailSent(){this.setProperties({_requestInProgress:!1,_password:""}),this.fire("cloud-done",{flashMessage:"Account created! Check your email for instructions on how to activate your account."})}});{const e=["/cloud/account"],t=["/cloud/login","/cloud/register","/cloud/forgot-password"];class i extends(window.hassMixins.NavigateMixin(s.a)){static get template(){return a["a"]`
    <app-route route="[[route]]" pattern="/cloud/:page" data="{{_routeData}}" tail="{{_routeTail}}"></app-route>

    <template is="dom-if" if="[[_equals(_routeData.page, &quot;account&quot;)]]" restamp="">
      <ha-config-cloud-account hass="[[hass]]" account="[[account]]" is-wide="[[isWide]]"></ha-config-cloud-account>
    </template>

    <template is="dom-if" if="[[_equals(_routeData.page, &quot;login&quot;)]]" restamp="">
      <ha-config-cloud-login page-name="login" hass="[[hass]]" is-wide="[[isWide]]" email="{{_loginEmail}}" flash-message="{{_flashMessage}}"></ha-config-cloud-login>
    </template>

    <template is="dom-if" if="[[_equals(_routeData.page, &quot;register&quot;)]]" restamp="">
      <ha-config-cloud-register page-name="register" hass="[[hass]]" is-wide="[[isWide]]" email="{{_loginEmail}}"></ha-config-cloud-register>
    </template>

    <template is="dom-if" if="[[_equals(_routeData.page, &quot;forgot-password&quot;)]]" restamp="">
      <ha-config-cloud-forgot-password page-name="forgot-password" hass="[[hass]]" email="{{_loginEmail}}"></ha-config-cloud-forgot-password>
    </template>
`}static get properties(){return{hass:Object,isWide:Boolean,loadingAccount:{type:Boolean,value:!1},account:{type:Object},_flashMessage:{type:String,value:""},route:Object,_routeData:Object,_routeTail:Object,_loginEmail:String}}static get observers(){return["_checkRoute(route, account)"]}ready(){super.ready(),this.addEventListener("cloud-done",e=>{this._flashMessage=e.detail.flashMessage,this.navigate("/config/cloud/login")})}_checkRoute(i){i&&"/cloud"===i.path.substr(0,6)&&(this._debouncer=Q.a.debounce(this._debouncer,K.timeOut.after(0),()=>{this.account||t.includes(i.path)?this.account&&!e.includes(i.path)&&this.navigate("/config/cloud/account",!0):this.navigate("/config/cloud/login",!0)}))}_equals(e,t){return e===t}}customElements.define("ha-config-cloud",i)}i(173),i(178),i(139);customElements.define("ha-config-flow",class extends(window.hassMixins.LocalizeMixin(window.hassMixins.EventsMixin(s.a))){static get template(){return a["a"]`
    <style include="ha-style-dialog">
      .error {
        color: red;
      }
      paper-dialog {
        max-width: 500px;
      }
      ha-markdown img:first-child:last-child {
        display: block;
        margin: 0 auto;
      }
    </style>
    <paper-dialog id="dialog" with-backdrop="" opened="[[step]]" on-opened-changed="_openedChanged">
      <h2>
        <template is="dom-if" if="[[_equals(step.type, &quot;abort&quot;)]]">
          Aborted
        </template>
        <template is="dom-if" if="[[_equals(step.type, &quot;create_entry&quot;)]]">
          Success!
        </template>
        <template is="dom-if" if="[[_equals(step.type, &quot;form&quot;)]]">
          [[_computeStepTitle(localize, step)]]
        </template>
      </h2>
      <paper-dialog-scrollable>
        <template is="dom-if" if="[[!step]]">
          Loading flow.
        </template>
        <template is="dom-if" if="[[step]]">
          <template is="dom-if" if="[[_equals(step.type, &quot;abort&quot;)]]">
            <p>[[_computeStepAbortedReason(localize, step)]]</p>
          </template>

          <template is="dom-if" if="[[_equals(step.type, &quot;create_entry&quot;)]]">
            <p>Created config for [[step.title]]</p>
          </template>

          <template is="dom-if" if="[[_equals(step.type, &quot;form&quot;)]]">
            <template is="dom-if" if="[[_computeStepDescription(localize, step)]]">
              <ha-markdown content="[[_computeStepDescription(localize, step)]]"></ha-markdown>
            </template>

            <ha-form data="{{stepData}}" schema="[[step.data_schema]]" error="[[step.errors]]" compute-label="[[_computeLabelCallback(localize, step)]]" compute-error="[[_computeErrorCallback(localize, step)]]"></ha-form>
          </template>
        </template>
      </paper-dialog-scrollable>
      <div class="buttons">
        <template is="dom-if" if="[[_equals(step.type, &quot;abort&quot;)]]">
          <paper-button on-click="_flowDone">Close</paper-button>
        </template>
        <template is="dom-if" if="[[_equals(step.type, &quot;create_entry&quot;)]]">
          <paper-button on-click="_flowDone">Close</paper-button>
        </template>
        <template is="dom-if" if="[[_equals(step.type, &quot;form&quot;)]]">
          <paper-button on-click="_submitStep">Submit</paper-button>
        </template>
      </div>
    </paper-dialog>
`}static get properties(){return{hass:Object,step:{type:Object,notify:!0},flowId:{type:String,observer:"_flowIdChanged"},stepData:Object}}ready(){super.ready(),this.addEventListener("keypress",e=>{13===e.keyCode&&this._submitStep()}),this.$.dialog.addEventListener("iron-overlay-opened",e=>{e.target.withBackdrop&&e.target.parentNode.insertBefore(e.target.backdropElement,e.target)})}_flowIdChanged(e){e?this.step?this._processStep(this.step):this.hass.callApi("get",`config/config_entries/flow/${e}`).then(e=>{this._processStep(e),setTimeout(()=>this.$.dialog.center(),0)}):this.setProperties({step:null,stepData:{}})}_submitStep(){this.hass.callApi("post",`config/config_entries/flow/${this.flowId}`,this.stepData).then(e=>this._processStep(e))}_processStep(e){e.errors||(e.errors={}),this.step=e,0===Object.keys(e.errors).length&&(this.stepData={})}_flowDone(){this.fire("flow-closed",{flowFinished:!0})}_equals(e,t){return e===t}_openedChanged(e){this.step&&!e.detail.value&&this.fire("flow-closed",{flowFinished:["success","abort"].includes(this.step.type)})}_computeStepAbortedReason(e,t){return e(`component.${t.handler}.config.abort.${t.reason}`)}_computeStepTitle(e,t){return e(`component.${t.handler}.config.step.${t.step_id}.title`)}_computeStepDescription(e,t){return e(`component.${t.handler}.config.step.${t.step_id}.description`)}_computeLabelCallback(e,t){return i=>e(`component.${t.handler}.config.step.${t.step_id}.data.${i.name}`)}_computeErrorCallback(e,t){return i=>e(`component.${t.handler}.config.error.${i}`)}});{class e extends(window.hassMixins.LocalizeMixin(window.hassMixins.EventsMixin(s.a))){static get template(){return a["a"]`
    <style include="iron-flex ha-style">
      paper-button {
        color: var(--primary-color);
        font-weight: 500;
        top: 3px;
        margin-right: -.57em;
      }
      paper-card:last-child {
        margin-top: 12px;
      }
      .config-entry-row {
        display: flex;
        padding: 0 16px;
      }
    </style>

    <hass-subpage header="Integrations">
      <template is="dom-if" if="[[_progress.length]]">
        <ha-config-section is-wide="[[isWide]]">
          <span slot="header">In Progress</span>
          <paper-card>
            <template is="dom-repeat" items="[[_progress]]">
              <div class="config-entry-row">
                <paper-item-body>
                  [[_computeIntegrationTitle(localize, item.handler)]]
                </paper-item-body>
                <paper-button on-click="_continueFlow">Configure</paper-button>
              </div>
            </template>
          </paper-card>
        </ha-config-section>
      </template>

      <ha-config-section is-wide="[[isWide]]">
        <span slot="header">Configured</span>
        <paper-card>
          <template is="dom-if" if="[[!_entries.length]]">
            <div class="config-entry-row">
              <paper-item-body>
                Nothing configured yet
              </paper-item-body>
            </div>
          </template>
          <template is="dom-repeat" items="[[_entries]]">
            <div class="config-entry-row">
              <paper-item-body three-line="">
                [[item.title]]
                <div secondary="">Integration: [[_computeIntegrationTitle(localize, item.domain)]]</div>
                <div secondary="">Added by: [[item.source]]</div>
                <div secondary="">State: [[item.state]]</div>
              </paper-item-body>
              <paper-button on-click="_removeEntry">Remove</paper-button>
            </div>
          </template>
        </paper-card>
      </ha-config-section>

      <ha-config-section is-wide="[[isWide]]">
        <span slot="header">Available</span>
        <paper-card>
          <template is="dom-repeat" items="[[_handlers]]">
            <div class="config-entry-row">
              <paper-item-body>
                [[_computeIntegrationTitle(localize, item)]]
              </paper-item-body>
              <paper-button on-click="_createFlow">Configure</paper-button>
            </div>
          </template>
        </paper-card>
      </ha-config-section>
    </hass-subpage>

    <ha-config-flow hass="[[hass]]" flow-id="[[_flowId]]" step="{{_flowStep}}" on-flow-closed="_flowClose"></ha-config-flow>
`}static get properties(){return{hass:Object,isWide:Boolean,_flowId:{type:String,value:null},_flowStep:Object,_entries:Array,_progress:Array,_handlers:Array}}ready(){super.ready(),this._loadData()}_createFlow(e){this.hass.callApi("post","config/config_entries/flow",{handler:e.model.item}).then(e=>{this._userCreatedFlow=!0,this.setProperties({_flowStep:e,_flowId:e.flow_id})})}_continueFlow(e){this._userCreatedFlow=!1,this.setProperties({_flowId:e.model.item.flow_id,_flowStep:null})}_removeEntry(e){if(!confirm("Are you sure you want to delete this integration?"))return;const t=e.model.item.entry_id;this.hass.callApi("delete",`config/config_entries/entry/${t}`).then(e=>{this._entries=this._entries.filter(e=>e.entry_id!==t),e.require_restart&&alert("Restart Home Assistant to finish removing this integration")})}_flowClose(e){e.detail.flowFinished?this._loadData():this._userCreatedFlow&&this.hass.callApi("delete",`config/config_entries/flow/${this._flowId}`),this._flowId=null}_loadData(){this._loadEntries(),this._loadDiscovery(),this.hass.callApi("get","config/config_entries/flow_handlers").then(e=>{this._handlers=e})}_loadEntries(){this.hass.callApi("get","config/config_entries/entry").then(e=>{this._entries=e})}_loadDiscovery(){this.hass.callApi("get","config/config_entries/flow").then(e=>{this._progress=e})}_computeIntegrationTitle(e,t){return e(`component.${t}.config.title`)}}customElements.define("ha-config-entries",e)}i(166);var ee=i(120);customElements.define("ha-config-section-core",class extends(window.hassMixins.LocalizeMixin(s.a)){static get template(){return a["a"]`
    <style include="iron-flex ha-style">
      .validate-container {
        @apply --layout-vertical;
        @apply --layout-center-center;
        height: 140px;
      }

      .validate-result {
        color: var(--google-green-500);
        font-weight: 500;
        margin-bottom: 1em;
      }

      .config-invalid {
        margin: 1em 0;
      }

      .config-invalid .text {
        color: var(--google-red-500);
        font-weight: 500;
      }

      .config-invalid paper-button {
        float: right;
      }

      .validate-log {
        white-space: pre-wrap;
      }
    </style>
    <ha-config-section is-wide="[[isWide]]">
      <span slot="header">[[localize('ui.panel.config.core.section.core.header')]]</span>
      <span slot="introduction">[[localize('ui.panel.config.core.section.core.introduction')]]</span>

      <paper-card heading="[[localize('ui.panel.config.core.section.core.validation.heading')]]">
        <div class="card-content">
          [[localize('ui.panel.config.core.section.core.validation.introduction')]]
          <template is="dom-if" if="[[!validateLog]]">
            <div class="validate-container">
              <template is="dom-if" if="[[!validating]]">
                <template is="dom-if" if="[[isValid]]">
                  <div class="validate-result" id="result">
                    [[localize('ui.panel.config.core.section.core.validation.valid')]]
                  </div>
                </template>
                <paper-button raised="" on-click="validateConfig">
                  [[localize('ui.panel.config.core.section.core.validation.check_config')]]
                </paper-button>
              </template>
              <template is="dom-if" if="[[validating]]">
                <paper-spinner active=""></paper-spinner>
              </template>
            </div>
          </template>
          <template is="dom-if" if="[[validateLog]]">
            <div class="config-invalid">
              <span class="text">
                [[localize('ui.panel.config.core.section.core.validation.invalid')]]
              </span>
              <paper-button raised="" on-click="validateConfig">
                [[localize('ui.panel.config.core.section.core.validation.check_config')]]
              </paper-button>
            </div>
            <div id="configLog" class="validate-log">[[validateLog]]</div>
          </template>
        </div>
      </paper-card>

      <paper-card heading="[[localize('ui.panel.config.core.section.core.reloading.heading')]]">
        <div class="card-content">
          [[localize('ui.panel.config.core.section.core.reloading.introduction')]]
        </div>
        <div class="card-actions">
          <ha-call-service-button hass="[[hass]]" domain="homeassistant" service="reload_core_config">[[localize('ui.panel.config.core.section.core.reloading.core')]]
          </ha-call-service-button>
          <ha-call-service-button hass="[[hass]]" domain="group" service="reload" hidden\$="[[!groupLoaded(hass)]]">[[localize('ui.panel.config.core.section.core.reloading.group')]]
          </ha-call-service-button>
          <ha-call-service-button hass="[[hass]]" domain="automation" service="reload" hidden\$="[[!automationLoaded(hass)]]">[[localize('ui.panel.config.core.section.core.reloading.automation')]]
          </ha-call-service-button>
          <ha-call-service-button hass="[[hass]]" domain="script" service="reload" hidden\$="[[!scriptLoaded(hass)]]">[[localize('ui.panel.config.core.section.core.reloading.script')]]
          </ha-call-service-button>
        </div>
      </paper-card>

      <paper-card heading="[[localize('ui.panel.config.core.section.core.server_management.heading')]]">
        <div class="card-content">
          [[localize('ui.panel.config.core.section.core.server_management.introduction')]]
        </div>
        <div class="card-actions warning">
          <ha-call-service-button class="warning" hass="[[hass]]" domain="homeassistant" service="restart">[[localize('ui.panel.config.core.section.core.server_management.restart')]]
          </ha-call-service-button>
          <ha-call-service-button class="warning" hass="[[hass]]" domain="homeassistant" service="stop">[[localize('ui.panel.config.core.section.core.server_management.stop')]]
          </ha-call-service-button>
        </div>
      </paper-card>

    </ha-config-section>
`}static get properties(){return{hass:{type:Object},isWide:{type:Boolean,value:!1},validating:{type:Boolean,value:!1},isValid:{type:Boolean,value:null},validateLog:{type:String,value:""}}}groupLoaded(e){return Object(ee.a)(e,"group")}automationLoaded(e){return Object(ee.a)(e,"automation")}scriptLoaded(e){return Object(ee.a)(e,"script")}validateConfig(){this.validating=!0,this.validateLog="",this.isValid=null,this.hass.callApi("POST","config/core/check_config").then(e=>{this.validating=!1,this.isValid="valid"===e.result,this.isValid||(this.validateLog=e.errors)})}});i(332),i(136);customElements.define("ha-push-notifications-toggle",class extends(window.hassMixins.EventsMixin(s.a)){static get template(){return a["a"]`
    <paper-toggle-button hidden\$="[[!pushSupported]]" disabled="[[loading]]" checked="{{pushChecked}}"></paper-toggle-button>
`}static get properties(){return{hass:{type:Object,value:null},pushSupported:{type:Boolean,readOnly:!0,notify:!0,value:"serviceWorker"in navigator&&"PushManager"in window&&("https:"===document.location.protocol||"localhost"===document.location.hostname||"127.0.0.1"===document.location.hostname)},pushChecked:{type:Boolean,value:"Notification"in window&&"granted"===Notification.permission,observer:"handlePushChange"},loading:{type:Boolean,value:!0}}}connectedCallback(){super.connectedCallback(),this.pushSupported&&navigator.serviceWorker.ready.then(e=>{e.pushManager.getSubscription().then(e=>{this.loading=!1,this.pushChecked=!!e})},()=>{this._setPushSupported(!1)})}handlePushChange(e){this.pushSupported&&(e?this.subscribePushNotifications():this.unsubscribePushNotifications())}subscribePushNotifications(){navigator.serviceWorker.ready.then(e=>e.pushManager.subscribe({userVisibleOnly:!0})).then(e=>{let t;return t=navigator.userAgent.toLowerCase().indexOf("firefox")>-1?"firefox":"chrome",this.hass.callApi("POST","notify.html5",{subscription:e,browser:t})},e=>{let t;t=e.message&&-1!==e.message.indexOf("gcm_sender_id")?"Please setup the notify.html5 platform.":"Notification registration failed.",console.error(e),this.fire("hass-notification",{message:t}),this.pushChecked=!1})}unsubscribePushNotifications(){navigator.serviceWorker.ready.then(e=>e.pushManager.getSubscription()).then(e=>e?this.hass.callApi("DELETE","notify.html5",{subscription:e}).then(()=>{e.unsubscribe()}):Promise.resolve()).catch(e=>{console.error("Error in unsub push",e),this.fire("hass-notification",{message:"Failed unsubscribing for push notifications."})})}});customElements.define("ha-config-section-push-notifications",class extends(window.hassMixins.LocalizeMixin(s.a)){static get template(){return a["a"]`
    <style include="iron-flex iron-flex-alignment iron-positioning">
      ha-push-notifications-toggle {
        margin-left: 16px;
      }
    </style>
    <ha-config-section is-wide="[[isWide]]">
      <span slot="header">[[localize('ui.panel.config.core.section.push_notifications.header')]]</span>
      <span slot="introduction">
        [[localize('ui.panel.config.core.section.push_notifications.introduction')]]
      </span>

      <paper-card>
        <div class="card-content">
          <iron-label class="horizontal layout">
            [[localize('ui.panel.config.core.section.push_notifications.push_notifications')]]
            <ha-push-notifications-toggle hass="[[hass]]" push-supported="{{pushSupported}}"></ha-push-notifications-toggle>
          </iron-label>
        </div>
      </paper-card>
    </ha-config-section>
`}static get properties(){return{hass:Object,isWide:Boolean,pushSupported:{type:Boolean,notify:!0}}}});i(97);customElements.define("ha-config-section-themes",class extends(window.hassMixins.LocalizeMixin(window.hassMixins.EventsMixin(s.a))){static get template(){return a["a"]`
    <ha-config-section is-wide="[[isWide]]">
      <span slot="header">[[localize('ui.panel.config.core.section.themes.header')]]</span>
      <span slot="introduction">
        [[localize('ui.panel.config.core.section.themes.introduction')]]
      </span>

      <paper-card>
        <div class="card-content">
          <paper-dropdown-menu label="[[localize('ui.panel.config.core.section.themes.header')]]" dynamic-align="">
            <paper-listbox slot="dropdown-content" selected="{{selectedTheme}}">
              <template is="dom-repeat" items="[[themes]]" as="theme">
                <paper-item>[[theme]]</paper-item>
              </template>
            </paper-listbox>
          </paper-dropdown-menu>
        </div>
      </paper-card>
    </ha-config-section>
`}static get properties(){return{hass:{type:Object},isWide:{type:Boolean},themes:{type:Array,computed:"computeThemes(hass)"},selectedTheme:{type:Number}}}static get observers(){return["selectionChanged(hass, selectedTheme)"]}ready(){super.ready(),this.hass.selectedTheme&&this.themes.indexOf(this.hass.selectedTheme)>0?this.selectedTheme=this.themes.indexOf(this.hass.selectedTheme):this.hass.selectedTheme||(this.selectedTheme=0)}computeThemes(e){return e?["Backend-selected","default"].concat(Object.keys(e.themes.themes).sort()):[]}selectionChanged(e,t){t>0&&t<this.themes.length?e.selectedTheme!==this.themes[t]&&this.fire("settheme",this.themes[t]):0===t&&""!==e.selectedTheme&&this.fire("settheme","")}});customElements.define("ha-config-section-translation",class extends(window.hassMixins.LocalizeMixin(window.hassMixins.EventsMixin(s.a))){static get template(){return a["a"]`
    <ha-config-section is-wide="[[isWide]]">
      <span slot="header">[[localize('ui.panel.config.core.section.translation.header')]]</span>
      <span slot="introduction">
        [[localize('ui.panel.config.core.section.translation.introduction')]]
      </span>

      <paper-card>
        <div class="card-content">
          <paper-dropdown-menu label="[[localize('ui.panel.config.core.section.translation.language')]]" dynamic-align="">
            <paper-listbox slot="dropdown-content" attr-for-selected="language-tag" selected="{{languageSelection}}">
              <template is="dom-repeat" items="[[languages]]">
                <paper-item language-tag\$="[[item.tag]]">[[item.nativeName]]</paper-item>
              </template>
            </paper-listbox>
          &gt;</paper-dropdown-menu>
        </div>
      </paper-card>
    </ha-config-section>
`}static get properties(){return{hass:{type:Object},isWide:{type:Boolean},languageSelection:{type:String,observer:"languageSelectionChanged"},languages:{type:Array,computed:"computeLanguages(hass)"}}}static get observers(){return["setLanguageSelection(language)"]}computeLanguages(e){return e&&e.translationMetadata?Object.keys(e.translationMetadata.translations).map(t=>({tag:t,nativeName:e.translationMetadata.translations[t].nativeName})):[]}setLanguageSelection(e){this.languageSelection=e}languageSelectionChanged(e){e!==this.language&&this.fire("hass-language-select",{language:e})}});customElements.define("ha-config-core",class extends(window.hassMixins.LocalizeMixin(s.a)){static get template(){return a["a"]`
    <style include="iron-flex ha-style">
      .content {
        padding-bottom: 32px;
      }

      .border {
        margin: 32px auto 0;
        border-bottom: 1px solid rgba(0, 0, 0, 0.12);
        max-width: 1040px;
      }

      .narrow .border {
        max-width: 640px;
      }
    </style>

    <ha-app-layout has-scrolling-region="">
      <app-header slot="header" fixed="">
        <app-toolbar>
          <paper-icon-button icon="mdi:arrow-left" on-click="_backTapped"></paper-icon-button>
          <div main-title="">[[localize('ui.panel.config.core.caption')]]</div>
        </app-toolbar>
      </app-header>

      <div class\$="[[computeClasses(isWide)]]">
        <ha-config-section-core is-wide="[[isWide]]" hass="[[hass]]"></ha-config-section-core>

        <template is="dom-if" if="[[pushSupported]]">
          <div class="border"></div>
          <ha-config-section-push-notifications is-wide="[[isWide]]" hass="[[hass]]" push-supported="{{pushSupported}}"></ha-config-section-push-notifications>
        </template>
        <template is="dom-if" if="[[computeIsTranslationLoaded(hass)]]">
          <div class="border"></div>
          <ha-config-section-translation is-wide="[[isWide]]" hass="[[hass]]"></ha-config-section-translation>
        </template>

        <template is="dom-if" if="[[computeIsThemesLoaded(hass)]]">
          <div class="border"></div>
          <ha-config-section-themes is-wide="[[isWide]]" hass="[[hass]]"></ha-config-section-themes>
        </template>

      </div>
    </ha-app-layout>
`}static get properties(){return{hass:Object,isWide:Boolean,pushSupported:{type:Boolean,value:!0}}}computeClasses(e){return e?"content":"content narrow"}computeIsZwaveLoaded(e){return Object(ee.a)(e,"config.zwave")}computeIsTranslationLoaded(e){return e.translationMetadata&&Object.keys(e.translationMetadata.translations).length}computeIsThemesLoaded(e){return e.themes&&e.themes.themes&&Object.keys(e.themes.themes).length}_backTapped(){history.back()}});i(99);customElements.define("ha-entity-config",class extends s.a{static get template(){return a["a"]`
    <style include="iron-flex ha-style">
      paper-card {
        display: block;
      }

      .device-picker {
        @apply --layout-horizontal;
        padding-bottom: 24px;
      }

      .form-placeholder {
        @apply --layout-vertical;
        @apply --layout-center-center;
        height: 96px;
      }

      [hidden]: {
        display: none;
      }

      .card-actions {
        @apply --layout-horizontal;
        @apply --layout-justified;
      }
    </style>
    <paper-card>
      <div class="card-content">
        <div class="device-picker">
          <paper-dropdown-menu label="[[label]]" class="flex" disabled="[[!entities.length]]">
            <paper-listbox slot="dropdown-content" selected="{{selectedEntity}}">
              <template is="dom-repeat" items="[[entities]]" as="state">
                <paper-item>[[computeSelectCaption(state)]]</paper-item>
              </template>
            </paper-listbox>
          </paper-dropdown-menu>
        </div>

        <div class="form-container">
          <template is="dom-if" if="[[computeShowPlaceholder(formState)]]">
            <div class="form-placeholder">
              <template is="dom-if" if="[[computeShowNoDevices(formState)]]">
                No entities found! :-(
              </template>

              <template is="dom-if" if="[[computeShowSpinner(formState)]]">
                <paper-spinner active="" alt="[[formState]]"></paper-spinner>
                [[formState]]
              </template>
            </div>
          </template>

          <div hidden\$="[[!computeShowForm(formState)]]" id="form"></div>
        </div>
      </div>
      <div class="card-actions">
        <paper-button on-click="saveEntity" disabled="[[computeShowPlaceholder(formState)]]">SAVE</paper-button>
        <template is="dom-if" if="[[allowDelete]]">
          <paper-button class="warning" on-click="deleteEntity" disabled="[[computeShowPlaceholder(formState)]]">DELETE</paper-button>
        </template>
      </div>
    </paper-card>
`}static get properties(){return{hass:{type:Object,observer:"hassChanged"},label:{type:String,value:"Device"},entities:{type:Array,observer:"entitiesChanged"},allowDelete:{type:Boolean,value:!1},selectedEntity:{type:Number,value:-1,observer:"entityChanged"},formState:{type:String,value:"no-devices"},config:{type:Object}}}connectedCallback(){super.connectedCallback(),this.formEl=document.createElement(this.config.component),this.formEl.hass=this.hass,this.$.form.appendChild(this.formEl),this.entityChanged(this.selectedEntity)}computeSelectCaption(e){return this.config.computeSelectCaption?this.config.computeSelectCaption(e):Object(Z.a)(e)}computeShowNoDevices(e){return"no-devices"===e}computeShowSpinner(e){return"loading"===e||"saving"===e}computeShowPlaceholder(e){return"editing"!==e}computeShowForm(e){return"editing"===e}hassChanged(e){this.formEl&&(this.formEl.hass=e)}entitiesChanged(e,t){if(0!==e.length)if(t){var i=t[this.selectedEntity].entity_id,a=e.findIndex(function(e){return e.entity_id===i});-1===a?this.selectedEntity=0:a!==this.selectedEntity&&(this.selectedEntity=a)}else this.selectedEntity=0;else this.formState="no-devices"}entityChanged(e){if(this.entities&&this.formEl){var t=this.entities[e];if(t){this.formState="loading";var i=this;this.formEl.loadEntity(t).then(function(){i.formState="editing"})}}}saveEntity(){this.formState="saving";var e=this;this.formEl.saveEntity().then(function(){e.formState="editing"})}});i(182);var te=i(30);i(194);customElements.define("ha-customize-array",class extends(window.hassMixins.EventsMixin(s.a)){static get template(){return a["a"]`
    <style>
      paper-dropdown-menu {
        margin: -9px 0;
      }
    </style>
    <paper-dropdown-menu label="[[item.description]]" disabled="[[item.secondary]]" selected-item-label="{{item.value}}" dynamic-align="">
      <paper-listbox slot="dropdown-content" selected="[[computeSelected(item)]]">
        <template is="dom-repeat" items="[[getOptions(item)]]" as="option">
          <paper-item>[[option]]</paper-item>
        </template>
      </paper-listbox>
    </paper-dropdown-menu>
`}static get properties(){return{item:{type:Object,notifies:!0}}}getOptions(e){const t=e.domain||"*",i=e.options[t]||e.options["*"];return i?i.sort():(this.item.type="string",this.fire("item-changed"),[])}computeSelected(e){return this.getOptions(e).indexOf(e.value)}});i(100);customElements.define("ha-customize-boolean",class extends s.a{static get template(){return a["a"]`
    <paper-checkbox disabled="[[item.secondary]]" checked="{{item.value}}">
      [[item.description]]
    </paper-checkbox>
`}static get properties(){return{item:{type:Object,notifies:!0}}}});i(43);customElements.define("ha-customize-icon",class extends s.a{static get template(){return a["a"]`
    <style>
      :host {
        @apply --layout-horizontal;
      }
      .icon-image {
        border: 1px solid grey;
        padding: 8px;
        margin-right: 20px;
        margin-top: 10px;
      }
    </style>
    <iron-icon class="icon-image" icon="[[item.value]]"></iron-icon>
    <paper-input auto-validate="" pattern="(mdi:.*)?" error-message="Must start with 'mdi:'" disabled="[[item.secondary]]" label="icon" value="{{item.value}}">
    </paper-input>
`}static get properties(){return{item:{type:Object,notifies:!0}}}});customElements.define("ha-customize-key-value",class extends s.a{static get template(){return a["a"]`
    <style>
      :host {
        @apply --layout-horizontal;
      }
      paper-input {
        @apply --layout-flex;
      }
      .key {
        padding-right: 20px;
      }
    </style>
    <paper-input disabled="[[item.secondary]]" class="key" label="Attribute name" value="{{item.attribute}}">
    </paper-input>
    <paper-input disabled="[[item.secondary]]" label="Attribute value" value="{{item.value}}">
    </paper-input>
`}static get properties(){return{item:{type:Object,notifies:!0}}}});customElements.define("ha-customize-string",class extends s.a{static get template(){return a["a"]`
    <paper-input disabled="[[item.secondary]]" label="[[getLabel(item)]]" value="{{item.value}}">
    </paper-input>
`}static get properties(){return{item:{type:Object,notifies:!0}}}getLabel(e){return e.description+("json"===e.type?" (JSON formatted)":"")}});customElements.define("ha-customize-attribute",class extends s.a{static get template(){return a["a"]`
    <style include="ha-form-style">
      :host {
        display: block;
        position: relative;
        padding-right: 40px;
      }

      .button {
        position: absolute;
        margin-top: -20px;
        top: 50%;
        right: 0;
      }
    </style>
    <div id="wrapper" class="form-group"></div>
    <paper-icon-button class="button" icon="[[getIcon(item.secondary)]]" on-click="tapButton"></paper-icon-button>
`}static get properties(){return{item:{type:Object,notify:!0,observer:"itemObserver"}}}tapButton(){this.item.secondary?this.item=Object.assign({},this.item,{secondary:!1}):this.item=Object.assign({},this.item,{closed:!0})}getIcon(e){return e?"mdi:pencil":"mdi:close"}itemObserver(e){const t=this.$.wrapper,i=window.hassAttributeUtil.TYPE_TO_TAG[e.type].toUpperCase();let a;t.lastChild&&t.lastChild.tagName===i?a=t.lastChild:(t.lastChild&&t.removeChild(t.lastChild),this.$.child=a=document.createElement(i.toLowerCase()),a.className="form-control",a.addEventListener("item-changed",()=>{this.item=Object.assign({},a.item)})),a.setProperties({item:this.item}),null===a.parentNode&&t.appendChild(a)}});customElements.define("ha-form-customize-attributes",class extends(Object(te.a)(s.a)){static get template(){return a["a"]`
    <style>
      [hidden] {
        display: none;
      }
    </style>
    <template is="dom-repeat" items="{{attributes}}" mutable-data="">
      <ha-customize-attribute item="{{item}}" hidden\$="[[item.closed]]">
      </ha-customize-attribute>
    </template>
`}static get properties(){return{attributes:{type:Array,notify:!0}}}});var ie=i(91);function ae(e,t){const i=Object(Z.a)(e),a=Object(Z.a)(t);return i<a?-1:i>a?1:0}customElements.define("ha-form-customize",class extends s.a{static get template(){return a["a"]`
    <style include="iron-flex ha-style ha-form-style">
      .warning {
        color: red;
      }

      .attributes-text {
        padding-left: 20px;
      }
    </style>
    <template is="dom-if" if="[[computeShowWarning(localConfig, globalConfig)]]">
      <div class="warning">
        It seems that your configuration.yaml doesn't properly include customize.yaml<br>
        Changes made here won't affect your configuration.
      </div>
    </template>
    <template is="dom-if" if="[[hasLocalAttributes]]">
      <h4 class="attributes-text">
        The following attributes are already set in customize.yaml<br>
      </h4>
      <ha-form-customize-attributes attributes="{{localAttributes}}"></ha-form-customize-attributes>
    </template>
    <template is="dom-if" if="[[hasGlobalAttributes]]">
      <h4 class="attributes-text">
        The following attributes are customized from outside of customize.yaml<br>
        Possibly via a domain, a glob or a different include.
      </h4>
      <ha-form-customize-attributes attributes="{{globalAttributes}}"></ha-form-customize-attributes>
    </template>
    <template is="dom-if" if="[[hasExistingAttributes]]">
      <h4 class="attributes-text">
        The following attributes of the entity are set programatically.<br>
        You can override them if you like.
      </h4>
      <ha-form-customize-attributes attributes="{{existingAttributes}}"></ha-form-customize-attributes>
    </template>
    <template is="dom-if" if="[[hasNewAttributes]]">
      <h4 class="attributes-text">
        The following attributes weren't set. Set them if you like.
      </h4>
      <ha-form-customize-attributes attributes="{{newAttributes}}"></ha-form-customize-attributes>
    </template>
    <div class="form-group">
      <paper-dropdown-menu label="Pick an attribute to override" class="flex" dynamic-align="">
        <paper-listbox slot="dropdown-content" selected="{{selectedNewAttribute}}">
          <template is="dom-repeat" items="[[newAttributesOptions]]" as="option">
            <paper-item>[[option]]</paper-item>
          </template>
        </paper-listbox>
      </paper-dropdown-menu>
    </div>
`}static get properties(){return{hass:{type:Object},entity:Object,localAttributes:{type:Array,computed:"computeLocalAttributes(localConfig)"},hasLocalAttributes:Boolean,globalAttributes:{type:Array,computed:"computeGlobalAttributes(localConfig, globalConfig)"},hasGlobalAttributes:Boolean,existingAttributes:{type:Array,computed:"computeExistingAttributes(localConfig, globalConfig, entity)"},hasExistingAttributes:Boolean,newAttributes:{type:Array,value:[]},hasNewAttributes:Boolean,newAttributesOptions:Array,selectedNewAttribute:{type:Number,value:-1,observer:"selectedNewAttributeObserver"},localConfig:Object,globalConfig:Object}}static get observers(){return["attributesObserver(localAttributes.*, globalAttributes.*, existingAttributes.*, newAttributes.*)"]}_initOpenObject(e,t,i,a){return Object.assign({attribute:e,value:t,closed:!1,domain:Object(ie.a)(this.entity),secondary:i,description:e},a)}loadEntity(e){return this.entity=e,this.hass.callApi("GET","config/customize/config/"+e.entity_id).then(e=>{this.localConfig=e.local,this.globalConfig=e.global,this.newAttributes=[]})}saveEntity(){const e={};this.localAttributes.concat(this.globalAttributes,this.existingAttributes,this.newAttributes).forEach(t=>{if(t.closed||t.secondary||!t.attribute||!t.value)return;const i="json"===t.type?JSON.parse(t.value):t.value;i&&(e[t.attribute]=i)});const t=this.entity.entity_id;return this.hass.callApi("POST","config/customize/config/"+t,e)}_computeSingleAttribute(e,t,i){const a=window.hassAttributeUtil.LOGIC_STATE_ATTRIBUTES[e]||{type:window.hassAttributeUtil.UNKNOWN_TYPE};return this._initOpenObject(e,"json"===a.type?JSON.stringify(t):t,i,a)}_computeAttributes(e,t,i){return t.map(t=>this._computeSingleAttribute(t,e[t],i))}computeLocalAttributes(e){if(!e)return[];const t=Object.keys(e);return this._computeAttributes(e,t,!1)}computeGlobalAttributes(e,t){if(!e||!t)return[];const i=Object.keys(e),a=Object.keys(t).filter(e=>!i.includes(e));return this._computeAttributes(t,a,!0)}computeExistingAttributes(e,t,i){if(!e||!t||!i)return[];const a=Object.keys(e),s=Object.keys(t),o=Object.keys(i.attributes).filter(e=>!a.includes(e)&&!s.includes(e));return this._computeAttributes(i.attributes,o,!0)}computeShowWarning(e,t){return!(!e||!t)&&Object.keys(e).some(i=>JSON.stringify(t[i])!==JSON.stringify(e[i]))}filterFromAttributes(e){return t=>!e||e.every(e=>e.attribute!==t||e.closed)}getNewAttributesOptions(e,t,i,a){return Object.keys(window.hassAttributeUtil.LOGIC_STATE_ATTRIBUTES).filter(e=>{const t=window.hassAttributeUtil.LOGIC_STATE_ATTRIBUTES[e];return t&&(!t.domains||!this.entity||t.domains.includes(Object(ie.a)(this.entity)))}).filter(this.filterFromAttributes(e)).filter(this.filterFromAttributes(t)).filter(this.filterFromAttributes(i)).filter(this.filterFromAttributes(a)).sort().concat("Other")}selectedNewAttributeObserver(e){if(e<0)return;const t=this.newAttributesOptions[e];if(e===this.newAttributesOptions.length-1){const e=this._initOpenObject("","",!1,{type:window.hassAttributeUtil.ADD_TYPE});return this.push("newAttributes",e),void(this.selectedNewAttribute=-1)}let i=this.localAttributes.findIndex(e=>e.attribute===t);if(i>=0)return this.set("localAttributes."+i+".closed",!1),void(this.selectedNewAttribute=-1);if((i=this.globalAttributes.findIndex(e=>e.attribute===t))>=0)return this.set("globalAttributes."+i+".closed",!1),void(this.selectedNewAttribute=-1);if((i=this.existingAttributes.findIndex(e=>e.attribute===t))>=0)return this.set("existingAttributes."+i+".closed",!1),void(this.selectedNewAttribute=-1);if((i=this.newAttributes.findIndex(e=>e.attribute===t))>=0)return this.set("newAttributes."+i+".closed",!1),void(this.selectedNewAttribute=-1);const a=this._computeSingleAttribute(t,"",!1);this.push("newAttributes",a),this.selectedNewAttribute=-1}attributesObserver(){this.hasLocalAttributes=this.localAttributes&&this.localAttributes.some(e=>!e.closed),this.hasGlobalAttributes=this.globalAttributes&&this.globalAttributes.some(e=>!e.closed),this.hasExistingAttributes=this.existingAttributes&&this.existingAttributes.some(e=>!e.closed),this.hasNewAttributes=this.newAttributes&&this.newAttributes.some(e=>!e.closed),this.newAttributesOptions=this.getNewAttributesOptions(this.localAttributes,this.globalAttributes,this.existingAttributes,this.newAttributes)}});customElements.define("ha-config-customize",class extends(window.hassMixins.LocalizeMixin(s.a)){static get template(){return a["a"]`
    <style include="ha-style">
    </style>

    <app-header-layout has-scrolling-region="">
      <app-header slot="header" fixed="">
        <app-toolbar>
          <paper-icon-button icon="mdi:arrow-left" on-click="_backTapped"></paper-icon-button>
          <div main-title="">[[localize('ui.panel.config.customize.caption')]]</div>
        </app-toolbar>
      </app-header>

      <div class\$="[[computeClasses(isWide)]]">
        <ha-config-section is-wide="[[isWide]]">
          <span slot="header">Customization</span>
          <span slot="introduction">
            Tweak per-entity attributes.<br>
            Added/edited customizations will take effect immediately. Removed customizations will take effect when the entity is updated.
          </span>
          <ha-entity-config hass="[[hass]]" label="Entity" entities="[[entities]]" config="[[entityConfig]]">
          </ha-entity-config>
        </ha-config-section>
      </div>
    </app-header-layout>
`}static get properties(){return{hass:Object,isWide:Boolean,entities:{type:Array,computed:"computeEntities(hass)"},entityConfig:{type:Object,value:{component:"ha-form-customize",computeSelectCaption:e=>Object(Z.a)(e)+" ("+Object(f.a)(e)+")"}}}}computeClasses(e){return e?"content":"content narrow"}_backTapped(){history.back()}computeEntities(e){return Object.keys(e.states).map(t=>e.states[t]).sort(ae)}});i(118);customElements.define("ha-config-cloud-menu",class extends(window.hassMixins.NavigateMixin(s.a)){static get template(){return a["a"]`
    <style include="iron-flex">
      paper-card {
        display: block;
      }
      paper-item {
        cursor: pointer;
      }
    </style>
    <paper-card>
      <paper-item on-click="_navigate">
        <paper-item-body two-line="">
          Home Assistant Cloud
          <template is="dom-if" if="[[account]]">
            <div secondary="">Logged in as [[account.email]]</div>
          </template>
          <template is="dom-if" if="[[!account]]">
            <div secondary="">Not logged in</div>
          </template>
        </paper-item-body>
        <iron-icon icon="mdi:chevron-right"></iron-icon>
      </paper-item>
    </paper-card>
`}static get properties(){return{hass:Object,isWide:Boolean,account:Object}}_navigate(){this.navigate("/config/cloud")}});customElements.define("ha-config-entries-menu",class extends(window.hassMixins.NavigateMixin(s.a)){static get template(){return a["a"]`
    <style include="iron-flex">
      paper-card {
        display: block;
      }
      paper-item {
        cursor: pointer;
      }
    </style>
    <paper-card>
      <paper-item on-click="_navigate">
        <paper-item-body two-line="">
          Integrations
          <div secondary="">EXPERIMENTAL – Manage connected devices and services</div>
        </paper-item-body>
        <iron-icon icon="mdi:chevron-right"></iron-icon>
      </paper-item>
    </paper-card>
`}static get properties(){return{hass:Object,isWide:Boolean,account:Object}}_navigate(){this.navigate("/config/integrations")}});{const e=["core","customize"];class t extends(window.hassMixins.LocalizeMixin(window.hassMixins.NavigateMixin(s.a))){static get template(){return a["a"]`
    <style include="iron-flex">
      paper-card {
        display: block;
      }
      paper-item {
        cursor: pointer;
      }
    </style>
    <paper-card>
      <template is="dom-repeat" items="[[pages]]">
        <template is="dom-if" if="[[_computeLoaded(hass, item)]]">
          <paper-item on-click="_navigate">
            <paper-item-body two-line="">
              [[_computeCaption(item, localize)]]
              <div secondary="">[[_computeDescription(item, localize)]]</div>
            </paper-item-body>
            <iron-icon icon="mdi:chevron-right"></iron-icon>
          </paper-item>
        </template>
      </template>
    </paper-card>
`}static get properties(){return{hass:{type:Object},pages:{type:Array,value:["core","customize","automation","script","zwave"]}}}_computeLoaded(t,i){return e.includes(i)||Object(ee.a)(t,i)}_computeCaption(e,t){return t(`ui.panel.config.${e}.caption`)}_computeDescription(e,t){return t(`ui.panel.config.${e}.description`)}_navigate(e){this.navigate("/config/"+e.model.item)}}customElements.define("ha-config-navigation",t)}customElements.define("ha-config-dashboard",class extends(window.hassMixins.LocalizeMixin(s.a)){static get template(){return a["a"]`
    <style include="iron-flex ha-style">
      .content {
        padding-bottom: 32px;
      }
    </style>

    <app-header-layout has-scrolling-region="">
      <app-header slot="header" fixed="">
        <app-toolbar>
          <ha-menu-button narrow="[[narrow]]" show-menu="[[showMenu]]"></ha-menu-button>
          <div main-title="">[[localize('panel.configuration')]]</div>
        </app-toolbar>
      </app-header>

      <div class="content">
        <ha-config-section is-wide="[[isWide]]">
          <span slot="header">[[localize('ui.panel.config.header')]]</span>
          <span slot="introduction">[[localize('ui.panel.config.introduction')]]</span>

          <template is="dom-if" if="[[computeIsLoaded(hass, &quot;cloud&quot;)]]">
            <ha-config-cloud-menu hass="[[hass]]" account="[[account]]"></ha-config-cloud-menu>
          </template>

          <template is="dom-if" if="[[computeIsLoaded(hass, &quot;config.config_entries&quot;)]]">
            <ha-config-entries-menu hass="[[hass]]"></ha-config-entries-menu>
          </template>

          <ha-config-navigation hass="[[hass]]"></ha-config-navigation>
        </ha-config-section>
      </div>
    </app-header-layout>
`}static get properties(){return{hass:Object,isWide:Boolean,account:Object,narrow:Boolean,showMenu:Boolean}}computeIsLoaded(e,t){return Object(ee.a)(e,t)}});class se extends o.a{constructor(){super(),this.onChange=this.onChange.bind(this),this.sequenceChanged=this.sequenceChanged.bind(this)}onChange(e){this.props.onChange(Object.assign({},this.props.script,{[e.target.name]:e.target.value}))}sequenceChanged(e){this.props.onChange(Object.assign({},this.props.script,{sequence:e}))}render({script:e,isWide:t,hass:i,localize:a}){const{alias:s,sequence:n}=e;return Object(o.b)("div",null,Object(o.b)("ha-config-section",{"is-wide":t},Object(o.b)("span",{slot:"header"},s),Object(o.b)("span",{slot:"introduction"},"Use scripts to execute a sequence of actions."),Object(o.b)("paper-card",null,Object(o.b)("div",{class:"card-content"},Object(o.b)("paper-input",{label:"Name",name:"alias",value:s,"onvalue-changed":this.onChange})))),Object(o.b)("ha-config-section",{"is-wide":t},Object(o.b)("span",{slot:"header"},"Sequence"),Object(o.b)("span",{slot:"introduction"},"The sequence of actions of this script.",Object(o.b)("p",null,Object(o.b)("a",{href:"https://home-assistant.io/docs/scripts/",target:"_blank"},"Learn more about available actions."))),Object(o.b)(F,{script:n,onChange:this.sequenceChanged,hass:i,localize:a})))}}var oe=i(125);customElements.define("ha-script-editor",class extends(window.hassMixins.LocalizeMixin(window.hassMixins.NavigateMixin(s.a))){static get template(){return a["a"]`
    <style include="ha-style">
      .errors {
        padding: 20px;
        font-weight: bold;
        color: var(--google-red-500);
      }
      .content {
        padding-bottom: 20px;
      }
      paper-card {
        display: block;
      }
      .triggers,
      .script {
        margin-top: -16px;
      }
      .triggers paper-card,
      .script paper-card {
        margin-top: 16px;
      }
      .add-card paper-button {
        display: block;
        text-align: center;
      }
      .card-menu {
        position: absolute;
        top: 0;
        right: 0;
        z-index: 1;
        color: var(--primary-text-color);
      }
      .card-menu paper-item {
        cursor: pointer;
      }
      span[slot=introduction] a {
        color: var(--primary-color);
      }
      paper-fab {
        position: fixed;
        bottom: 16px;
        right: 16px;
        z-index: 1;
        margin-bottom: -80px;
        transition: margin-bottom .3s;
      }

      paper-fab[is-wide] {
        bottom: 24px;
        right: 24px;
      }

      paper-fab[dirty] {
        margin-bottom: 0;
      }
    </style>
    <ha-app-layout has-scrolling-region="">
      <app-header slot="header" fixed="">
        <app-toolbar>
          <paper-icon-button icon="mdi:arrow-left" on-click="backTapped"></paper-icon-button>
          <div main-title="">Script [[name]]</div>
        </app-toolbar>
      </app-header>
      <div class="content">
        <template is="dom-if" if="[[errors]]">
          <div class="errors">[[errors]]</div>
        </template>
        <div id="root"></div>
      </div>
      <paper-fab slot="fab" is-wide\$="[[isWide]]" dirty\$="[[dirty]]" icon="mdi:content-save" title="Save" on-click="saveScript"></paper-fab>
    </ha-app-layout>
`}static get properties(){return{hass:{type:Object},narrow:{type:Boolean},showMenu:{type:Boolean,value:!1},errors:{type:Object,value:null},dirty:{type:Boolean,value:!1},config:{type:Object,value:null},script:{type:Object,observer:"scriptChanged"},creatingNew:{type:Boolean,observer:"creatingNewChanged"},name:{type:String,computed:"computeName(script)"},isWide:{type:Boolean,observer:"_updateComponent"},_rendered:{type:Object,value:null},_renderScheduled:{type:Boolean,value:!1}}}ready(){this.configChanged=this.configChanged.bind(this),super.ready()}disconnectedCallback(){super.disconnectedCallback(),this._rendered&&(Y(this._rendered),this._rendered=null)}configChanged(e){null!==this._rendered&&(this.config=e,this.errors=null,this.dirty=!0,this._updateComponent())}scriptChanged(e,t){e&&(this.hass?t&&t.entity_id===e.entity_id||this.hass.callApi("get","config/script/config/"+Object(oe.a)(e.entity_id)).then(e=>{var t=e.sequence;t&&!Array.isArray(t)&&(e.sequence=[t]),this.dirty=!1,this.config=e,this._updateComponent()},()=>{alert("Only scripts inside scripts.yaml are editable."),history.back()}):setTimeout(()=>this.scriptChanged(e,t),0))}creatingNewChanged(e){e&&(this.dirty=!1,this.config={alias:"New Script",sequence:[{service:"",data:{}}]},this._updateComponent())}backTapped(){this.dirty&&!confirm("You have unsaved changes. Are you sure you want to leave?")||history.back()}_updateComponent(){var e,t,i;!this._renderScheduled&&this.hass&&this.config&&(this._renderScheduled=!0,Promise.resolve().then(()=>{this._rendered=(e=this.$.root,t={script:this.config,onChange:this.configChanged,isWide:this.isWide,hass:this.hass,localize:this.localize},i=this._rendered,Object(o.c)(Object(o.b)(se,t),e,i)),this._renderScheduled=!1}))}saveScript(){var e=this.creatingNew?""+Date.now():Object(oe.a)(this.script.entity_id);this.hass.callApi("post","config/script/config/"+e,this.config).then(()=>{this.dirty=!1,this.creatingNew&&this.navigate(`/config/script/edit/${e}`,!0)},e=>{throw this.errors=e.body.message,e})}computeName(e){return e&&Object(Z.a)(e)}});customElements.define("ha-script-picker",class extends(window.hassMixins.LocalizeMixin(window.hassMixins.NavigateMixin(s.a))){static get template(){return a["a"]`
    <style include="ha-style">
      :host {
        display: block;
      }

      paper-item {
        cursor: pointer;
      }

      paper-fab {
        position: fixed;
        bottom: 16px;
        right: 16px;
        z-index: 1;
      }

      paper-fab[is-wide] {
        bottom: 24px;
        right: 24px;
      }

      a {
        color: var(--primary-color);
      }
    </style>

    <ha-app-layout has-scrolling-region="">
      <app-header slot="header" fixed="">
        <app-toolbar>
          <paper-icon-button icon="mdi:arrow-left" on-click="_backTapped"></paper-icon-button>
          <div main-title="">[[localize('ui.panel.config.script.caption')]]</div>
        </app-toolbar>
      </app-header>

      <ha-config-section is-wide="[[isWide]]">
        <div slot="header">Script Editor</div>
        <div slot="introduction">
          The script editor allows you to create and edit scripts.
          Please read <a href="https://home-assistant.io/docs/scripts/editor/" target="_blank">the instructions</a> to make sure that you have configured Home Assistant correctly.
        </div>

        <paper-card heading="Pick script to edit">
          <template is="dom-if" if="[[!scripts.length]]">
            <div class="card-content">
              <p>We couldn't find any editable scripts.</p>
            </div>
          </template>
          <template is="dom-repeat" items="[[scripts]]" as="script">
            <paper-item>
              <paper-item-body two-line="" on-click="scriptTapped">
                <div>[[computeName(script)]]</div>
                <div secondary="">[[computeDescription(script)]]</div>
              </paper-item-body>
              <iron-icon icon="mdi:chevron-right"></iron-icon>
            </paper-item>
          </template>
        </paper-card>
      </ha-config-section>

      <paper-fab slot="fab" is-wide\$="[[isWide]]" icon="mdi:plus" title="Add Script" on-click="addScript"></paper-fab>
    </ha-app-layout>
`}static get properties(){return{hass:{type:Object},narrow:{type:Boolean},showMenu:{type:Boolean,value:!1},scripts:{type:Array},isWide:{type:Boolean}}}scriptTapped(e){this.navigate("/config/script/edit/"+this.scripts[e.model.index].entity_id)}addScript(){this.navigate("/config/script/new")}computeName(e){return Object(Z.a)(e)}computeDescription(e){return""}_backTapped(){history.back()}});customElements.define("ha-config-script",class extends s.a{static get template(){return a["a"]`
    <style>
      ha-script-picker,
      ha-script-editor {
        height: 100%;
      }
    </style>
    <app-route route="[[route]]" pattern="/script/edit/:script" data="{{_routeData}}" active="{{_edittingScript}}"></app-route>
    <app-route route="[[route]]" pattern="/script/new" active="{{_creatingNew}}"></app-route>

    <template is="dom-if" if="[[!showEditor]]">
      <ha-script-picker hass="[[hass]]" narrow="[[narrow]]" show-menu="[[showMenu]]" scripts="[[scripts]]" is-wide="[[isWide]]"></ha-script-picker>
    </template>

    <template is="dom-if" if="[[showEditor]]" restamp="">
      <ha-script-editor hass="[[hass]]" script="[[script]]" is-wide="[[isWide]]" creating-new="[[_creatingNew]]"></ha-script-editor>
    </template>
`}static get properties(){return{hass:Object,narrow:Boolean,showMenu:Boolean,route:Object,isWide:Boolean,_routeData:Object,_routeMatches:Boolean,_creatingNew:Boolean,_edittingScript:Boolean,scripts:{type:Array,computed:"computeScripts(hass)"},script:{type:Object,computed:"computeScript(scripts, _edittingScript, _routeData)"},showEditor:{type:Boolean,computed:"computeShowEditor(_edittingScript, _creatingNew)"}}}computeScript(e,t,i){if(!e||!t)return null;for(var a=0;a<e.length;a++)if(e[a].entity_id===i.script)return e[a];return null}computeScripts(e){var t=[];return Object.keys(e.states).forEach(function(i){var a=e.states[i];"script"===Object(f.a)(a)&&t.push(a)}),t.sort(function(e,t){var i=Object(Z.a)(e),a=Object(Z.a)(t);return i<a?-1:i>a?1:0})}computeShowEditor(e,t){return t||e}});customElements.define("ha-service-description",class extends s.a{static get template(){return a["a"]`
    [[_getDescription(hass, domain, service)]]
`}static get properties(){return{hass:Object,domain:String,service:String}}_getDescription(e,t,i){var a=e.config.services[t];if(!a)return"";var s=a[i];return s?s.description:""}});customElements.define("zwave-groups",class extends s.a{static get template(){return a["a"]`
    <style include="iron-flex ha-style">
      .content {
        margin-top: 24px;
      }

      paper-card {
        display: block;
        margin: 0 auto;
        max-width: 600px;
      }

      .device-picker {
        @apply --layout-horizontal;
        @apply --layout-center-center;
        padding-left: 24px;
        padding-right: 24px;
        padding-bottom: 24px;
        }

      .help-text {
        padding-left: 24px;
        padding-right: 24px;
        padding-bottom: 12px;
      }
    </style>
    <paper-card class="content" heading="Node group associations">
      <!--TODO make api for getting groups and members-->
      <div class="device-picker">
        <paper-dropdown-menu label="Group" dynamic-align="" class="flex">
          <paper-listbox slot="dropdown-content" selected="{{selectedGroup}}">
            <template is="dom-repeat" items="[[groups]]" as="state">
              <paper-item>[[computeSelectCaptionGroup(state)]]</paper-item>
            </template>
          </paper-listbox>
        </paper-dropdown-menu>
      </div>
      <template is="dom-if" if="[[computeIsGroupSelected(selectedGroup)]]">
        <div class="device-picker">
          <paper-dropdown-menu label="Node to control" dynamic-align="" class="flex">
            <paper-listbox slot="dropdown-content" selected="{{selectedTargetNode}}">
              <template is="dom-repeat" items="[[nodes]]" as="state">
                <paper-item>[[computeSelectCaption(state)]]</paper-item>
              </template>
            </paper-listbox>
          </paper-dropdown-menu>
        </div>

        <div class="help-text">
          <span>Other Nodes in this group:</span>
          <template is="dom-repeat" items="[[otherGroupNodes]]" as="state">
            <div>[[state]]</div>
          </template>
        </div>
        <div class="help-text">
          <span>Max Associations:</span>
          <span>[[maxAssociations]]</span>
        </div>
      </template>

      <template is="dom-if" if="[[computeIsTargetNodeSelected(selectedTargetNode)]]">
        <div class="card-actions">
          <template is="dom-if" if="[[!noAssociationsLeft]]">
            <ha-call-service-button hass="[[hass]]" domain="zwave" service="change_association" service-data="[[computeAssocServiceData(selectedGroup, &quot;add&quot;)]]">Add To Group</ha-call-service-button>
          </template>
          <template is="dom-if" if="[[computeTargetInGroup(selectedGroup, selectedTargetNode)]]">
            <ha-call-service-button hass="[[hass]]" domain="zwave" service="change_association" service-data="[[computeAssocServiceData(selectedGroup, &quot;remove&quot;)]]">Remove From Group</ha-call-service-button>
          </template>
        </div>
      </template>
    </paper-card>
`}static get properties(){return{hass:{type:Object},nodes:{type:Array},groups:{type:Array},selectedNode:{type:Number},selectedTargetNode:{type:Number,value:-1},selectedGroup:{type:Number,value:-1,observer:"selectedGroupChanged"},otherGroupNodes:{type:Array,value:-1,computed:"computeOtherGroupNodes(selectedGroup)"},maxAssociations:{type:String,value:"",computed:"computeMaxAssociations(selectedGroup)"},noAssociationsLeft:{type:Boolean,value:!0,computed:"computeAssociationsLeft(selectedGroup)"}}}ready(){super.ready(),this.addEventListener("hass-service-called",e=>this.serviceCalled(e))}serviceCalled(e){if(e.detail.success){var t=this;setTimeout(function(){t.refreshGroups(t.selectedNode)},5e3)}}computeAssociationsLeft(e){return-1===e||this.maxAssociations===this.otherGroupNodes.length}computeMaxAssociations(e){if(-1===e)return-1;var t=this.groups[e].value.max_associations;return t||"None"}computeOtherGroupNodes(e){if(-1===e)return-1;var t=Object.values(this.groups[e].value.association_instances);return t.length?t.map(e=>{if(!e.length||2!==e.length)return"Unknown Node: "+e;const t=e[0],i=e[1],a=this.nodes.find(e=>e.attributes.node_id===t);if(!a)return"Unknown Node (id: "+(i?t+"."+i:t)+")";let s=this.computeSelectCaption(a);return i&&(s+="/ Instance: "+i),s}):["None"]}computeTargetInGroup(e,t){if(-1===e||-1===t)return!1;const i=Object.values(this.groups[e].value.associations);return!!i.length&&-1!==i.indexOf(this.nodes[t].attributes.node_id)}computeSelectCaption(e){return Object(Z.a)(e)+" (Node:"+e.attributes.node_id+" "+e.attributes.query_stage+")"}computeSelectCaptionGroup(e){return e.key+": "+e.value.label}computeIsTargetNodeSelected(e){return this.nodes&&-1!==e}computeIsGroupSelected(e){return this.nodes&&-1!==this.selectedNode&&-1!==e}computeAssocServiceData(e,t){return-1===!this.groups||-1===e||-1===this.selectedNode?-1:{node_id:this.nodes[this.selectedNode].attributes.node_id,association:t,target_node_id:this.nodes[this.selectedTargetNode].attributes.node_id,group:this.groups[e].key}}refreshGroups(e){var t=[];this.hass.callApi("GET","zwave/groups/"+this.nodes[e].attributes.node_id).then(function(e){Object.keys(e).forEach(function(i){t.push({key:i,value:e[i]})}),this.groups=t,this.selectedGroupChanged(this.selectedGroup)}.bind(this))}selectedGroupChanged(e){-1!==this.selectedGroup&&-1!==e&&(this.maxAssociations=this.groups[e].value.max_associations,this.otherGroupNodes=Object.values(this.groups[e].value.associations))}});customElements.define("ozw-log",class extends s.a{static get template(){return a["a"]`
    <style include="iron-flex ha-style">
      .content {
        margin-top: 24px;
      }

      paper-card {
        display: block;
        margin: 0 auto;
        max-width: 600px;
      }

      .device-picker {
        padding-left: 24px;
        padding-right: 24px;
        padding-bottom: 24px;
      }
    </style>
    <ha-config-section is-wide="[[isWide]]">
      <span slot="header">OZW Log</span>
      <paper-card>
      <div class="device-picker">
        <paper-input label="Number of last log lines." type="number" min="0" max="1000" step="10" value="{{numLogLines}}">
        </paper-input>
      </div>
      <div class="card-actions">
        <paper-button raised="" on-click="refreshLog">Refresh</paper-button>
      </div>
      <div class="help-text">
             <pre>[[ozwLogs]]</pre>
      </div>
      </paper-card>
    </ha-config-section>
`}static get properties(){return{hass:{type:Object},isWide:{type:Boolean,value:!1},ozwLogs:{type:String,value:"Refresh to pull log"},numLogLines:{type:Number,value:0}}}refreshLog(){this.ozwLogs="Loading ozw log...",this.hass.callApi("GET","zwave/ozwlog?lines="+this.numLogLines).then(e=>{this.ozwLogs=e})}});customElements.define("zwave-network",class extends s.a{static get template(){return a["a"]`
    <style include="iron-flex ha-style">
      .content {
        margin-top: 24px;
      }

      paper-card {
        display: block;
        margin: 0 auto;
        max-width: 600px;
      }

      .card-actions.warning ha-call-service-button {
        color: var(--google-red-500);
      }

      .toggle-help-icon {
        position: absolute;
        top: -6px;
        right: 0;
        color: var(--primary-color);
      }

      ha-service-description {
        display: block;
        color: grey;
      }

      [hidden] {
        display: none;
      }
    </style>
    <ha-config-section is-wide="[[isWide]]">
      <div style="position: relative" slot="header">
        <span>Z-Wave Network Management</span>
        <paper-icon-button class="toggle-help-icon" on-click="helpTap" icon="mdi:help-circle"></paper-icon-button>

      </div>
      <span slot="introduction">
        Run commands that affect the Z-Wave network. You won't get feedback on whether the command succeeded, but you can look in the OZW Log to try to figure out.
      </span>


      <paper-card class="content">
        <div class="card-actions">
          <ha-call-service-button hass="[[hass]]" domain="zwave" service="add_node_secure">Add Node Secure</ha-call-service-button>
          <ha-service-description hass="[[hass]]" domain="zwave" service="add_node_secure" hidden\$="[[!showDescription]]"></ha-service-description>

          <ha-call-service-button hass="[[hass]]" domain="zwave" service="add_node">Add Node</ha-call-service-button>
          <ha-service-description hass="[[hass]]" domain="zwave" service="add_node" hidden\$="[[!showDescription]]"></ha-service-description>

          <ha-call-service-button hass="[[hass]]" domain="zwave" service="remove_node">Remove Node</ha-call-service-button>
          <ha-service-description hass="[[hass]]" domain="zwave" service="remove_node" hidden\$="[[!showDescription]]"></ha-service-description>

        </div>
        <div class="card-actions warning">
          <ha-call-service-button hass="[[hass]]" domain="zwave" service="cancel_command">Cancel Command</ha-call-service-button>
          <ha-service-description hass="[[hass]]" domain="zwave" service="cancel_command" hidden\$="[[!showDescription]]"></ha-service-description>

        </div>
        <div class="card-actions">
          <ha-call-service-button hass="[[hass]]" domain="zwave" service="heal_network">Heal Network</ha-call-service-button>

          <ha-call-service-button hass="[[hass]]" domain="zwave" service="start_network">Start Network</ha-call-service-button>
          <ha-service-description hass="[[hass]]" domain="zwave" service="start_network" hidden\$="[[!showDescription]]"></ha-service-description>

          <ha-call-service-button hass="[[hass]]" domain="zwave" service="stop_network">Stop Network</ha-call-service-button>
          <ha-service-description hass="[[hass]]" domain="zwave" service="stop_network" hidden\$="[[!showDescription]]"></ha-service-description>

          <ha-call-service-button hass="[[hass]]" domain="zwave" service="soft_reset">Soft Reset</ha-call-service-button>
          <ha-service-description hass="[[hass]]" domain="zwave" service="soft_reset" hidden\$="[[!showDescription]]"></ha-service-description>

          <ha-call-service-button hass="[[hass]]" domain="zwave" service="test_network">Test Network</ha-call-service-button>
          <ha-service-description hass="[[hass]]" domain="zwave" service="test_network" hidden\$="[[!showDescription]]"></ha-service-description>
          <ha-call-api-button hass="[[hass]]" path="zwave/saveconfig">Save Config</ha-call-api-button>

        </div>
      </paper-card>
    </ha-config-section>
`}static get properties(){return{hass:{type:Object},isWide:{type:Boolean,value:!1},showDescription:{type:Boolean,value:!1}}}helpTap(){this.showDescription=!this.showDescription}});customElements.define("zwave-node-config",class extends s.a{static get template(){return a["a"]`
    <style include="iron-flex ha-style">
      .content {
        margin-top: 24px;
      }

      paper-card {
        display: block;
        margin: 0 auto;
        max-width: 600px;
      }

      .device-picker {
        @apply --layout-horizontal;
        @apply --layout-center-center;
        padding-left: 24px;
        padding-right: 24px;
        padding-bottom: 24px;
        }

      .help-text {
        padding-left: 24px;
        padding-right: 24px;
      }
    </style>
    <div class="content">
      <paper-card heading="Node config options">
        <template is="dom-if" if="[[wakeupNode]]">
          <div class="card-actions">
            <paper-input float-label="Wakeup Interval" type="number" value="{{wakeupInput}}" placeholder="[[computeGetWakeupValue(selectedNode)]]">
               <div suffix="">seconds</div>
            </paper-input>
            <ha-call-service-button hass="[[hass]]" domain="zwave" service="set_wakeup" service-data="[[computeWakeupServiceData(wakeupInput)]]">Set Wakeup</ha-call-service-button>
          </div>
        </template>
        <div class="device-picker">
        <paper-dropdown-menu label="Config parameter" dynamic-align="" class="flex">
          <paper-listbox slot="dropdown-content" selected="{{selectedConfigParameter}}">
            <template is="dom-repeat" items="[[config]]" as="state">
              <paper-item>[[computeSelectCaptionConfigParameter(state)]]</paper-item>
            </template>
          </paper-listbox>
        </paper-dropdown-menu>
        </div>
        <template is="dom-if" if="[[isConfigParameterSelected(selectedConfigParameter, 'List')]]">
          <div class="device-picker">
            <paper-dropdown-menu label="Config value" dynamic-align="" class="flex" placeholder="{{loadedConfigValue}}">
              <paper-listbox slot="dropdown-content" selected="{{selectedConfigValue}}">
                <template is="dom-repeat" items="[[selectedConfigParameterValues]]" as="state">
                  <paper-item>[[state]]</paper-item>
                </template>
              </paper-listbox>
            </paper-dropdown-menu>
          </div>
        </template>

        <template is="dom-if" if="[[isConfigParameterSelected(selectedConfigParameter, 'Byte Short Int')]]">
          <div class="card-actions">
            <paper-input label="{{selectedConfigParameterNumValues}}" type="number" value="{{selectedConfigValue}}" max="{{configParameterMax}}" min="{{configParameterMin}}">
            </paper-input>
          </div>
        </template>
        <template is="dom-if" if="[[isConfigParameterSelected(selectedConfigParameter, 'Bool Button')]]">
          <div class="device-picker">
            <paper-dropdown-menu label="Config value" class="flex" dynamic-align="" placeholder="{{loadedConfigValue}}">
              <paper-listbox slot="dropdown-content" selected="{{selectedConfigValue}}">
                <template is="dom-repeat" items="[[selectedConfigParameterValues]]" as="state">
                  <paper-item>[[state]]</paper-item>
                </template>
              </paper-listbox>
            </paper-dropdown-menu>
          </div>
        </template>
        <div class="help-text">
          <span>[[configValueHelpText]]</span>
        </div>
        <template is="dom-if" if="[[isConfigParameterSelected(selectedConfigParameter, 'Bool Button Byte Short Int List')]]">
          <div class="card-actions">
            <ha-call-service-button hass="[[hass]]" domain="zwave" service="set_config_parameter" service-data="[[computeSetConfigParameterServiceData(selectedConfigValue)]]">Set Config Parameter</ha-call-service-button>
          </div>
        </template>
      </paper-card>
    </div>
`}static get properties(){return{hass:{type:Object},nodes:{type:Array,observer:"nodesChanged"},selectedNode:{type:Number,value:-1,observer:"nodesChanged"},config:{type:Array,value:function(){return[]}},selectedConfigParameter:{type:Number,value:-1,observer:"selectedConfigParameterChanged"},configParameterMax:{type:Number,value:-1},configParameterMin:{type:Number,value:-1},configValueHelpText:{type:String,value:"",computed:"computeConfigValueHelp(selectedConfigParameter)"},selectedConfigParameterType:{type:String,value:""},selectedConfigValue:{type:Number,value:-1,observer:"computeSetConfigParameterServiceData"},selectedConfigParameterValues:{type:Array,value:function(){return[]}},selectedConfigParameterNumValues:{type:String,value:""},loadedConfigValue:{type:Number,value:-1},wakeupInput:{type:Number},wakeupNode:{type:Boolean,value:!1}}}ready(){super.ready(),this.addEventListener("hass-service-called",e=>this.serviceCalled(e))}serviceCalled(e){if(e.detail.success){var t=this;setTimeout(function(){t.refreshConfig(t.selectedNode)},5e3)}}nodesChanged(){this.nodes&&(this.wakeupNode=0===this.nodes[this.selectedNode].attributes.wake_up_interval||this.nodes[this.selectedNode].attributes.wake_up_interval,this.wakeupNode&&(0===this.nodes[this.selectedNode].attributes.wake_up_interval?this.wakeupInput="":this.wakeupInput=this.nodes[this.selectedNode].attributes.wake_up_interval))}computeGetWakeupValue(e){return-1!==this.selectedNode&&this.nodes[e].attributes.wake_up_interval?this.nodes[e].attributes.wake_up_interval:"unknown"}computeWakeupServiceData(e){return{node_id:this.nodes[this.selectedNode].attributes.node_id,value:e}}computeConfigValueHelp(e){if(-1===e)return"";var t=this.config[e].value.help;return t||["No helptext available"]}computeSetConfigParameterServiceData(e){if(-1===this.selectedNode||-1===this.selectedConfigParameter)return-1;var t=null;return"Short Byte Int".includes(this.selectedConfigParameterType)&&(t=parseInt(e,10)),"Bool Button".includes(this.selectedConfigParameterType)&&(t=this.selectedConfigParameterValues[e]),"List"===this.selectedConfigParameterType&&(t=this.selectedConfigParameterValues[e]),{node_id:this.nodes[this.selectedNode].attributes.node_id,parameter:this.config[this.selectedConfigParameter].key,value:t}}selectedConfigParameterChanged(e){-1!==e&&(this.selectedConfigValue=-1,this.loadedConfigValue=-1,this.selectedConfigParameterValues=[],this.selectedConfigParameterType=this.config[e].value.type,this.configParameterMax=this.config[e].value.max,this.configParameterMin=this.config[e].value.min,this.loadedConfigValue=this.config[e].value.data,this.configValueHelpText=this.config[e].value.help,"Short Byte Int".includes(this.selectedConfigParameterType)&&(this.selectedConfigParameterNumValues=this.config[e].value.data_items,this.selectedConfigValue=this.loadedConfigValue),"Bool Button".includes(this.selectedConfigParameterType)&&(this.selectedConfigParameterValues=["True","False"],this.config[e].value.data?this.loadedConfigValue="True":this.loadedConfigValue="False"),"List".includes(this.selectedConfigParameterType)&&(this.selectedConfigParameterValues=this.config[e].value.data_items))}isConfigParameterSelected(e,t){return-1!==e&&(this.config[e].value.type===t||!!t.includes(this.config[e].value.type))}computeSelectCaptionConfigParameter(e){return e.key+": "+e.value.label}refreshConfig(e){var t=[];this.hass.callApi("GET","zwave/config/"+this.nodes[e].attributes.node_id).then(function(e){Object.keys(e).forEach(function(i){t.push({key:i,value:e[i]})}),this.config=t,this.selectedConfigParameterChanged(this.selectedConfigParameter)}.bind(this))}});customElements.define("zwave-node-information",class extends s.a{static get template(){return a["a"]`
    <style include="iron-flex ha-style">
      .content {
        margin-top: 24px;
      }

      .node-info {
        margin-left: 16px;
      }

      paper-card {
        display: block;
        margin: 0 auto;
        max-width: 600px;
      }

      paper-button[toggles][active] {
        background: lightgray;
      }

    </style>

    <div class="content">
      <paper-card heading="Node Information">
      <div class="card-actions">
        <paper-button toggles="" raised="" noink="" active="{{nodeInfoActive}}">Show</paper-button>
      </div>
      <template is="dom-if" if="{{nodeInfoActive}}">
        <template is="dom-repeat" items="[[selectedNodeAttrs]]" as="state">
          <div class="node-info">
            <span>[[state]]</span>
          </div>
        </template>
      </template>
      </paper-card>
    </div>
`}static get properties(){return{nodes:{type:Array,observer:"nodeChanged"},selectedNode:{type:Number,value:-1,observer:"nodeChanged"},selectedNodeAttrs:{type:Array},nodeInfoActive:{type:Boolean}}}nodeChanged(e){if(this.nodes&&-1!==e){var t=this.nodes[this.selectedNode].attributes,i=[];Object.keys(t).forEach(function(e){i.push(e+": "+t[e])}),this.selectedNodeAttrs=i.sort()}}});customElements.define("zwave-usercodes",class extends s.a{static get template(){return a["a"]`
    <style include="iron-flex ha-style">
      .content {
        margin-top: 24px;
      }

      paper-card {
        display: block;
        margin: 0 auto;
        max-width: 600px;
      }

      .device-picker {
        @apply --layout-horizontal;
        @apply --layout-center-center;
        padding-left: 24px;
        padding-right: 24px;
        padding-bottom: 24px;
        }
    </style>
      <div class="content">
        <paper-card heading="Node user codes">
          <div class="device-picker">
          <paper-dropdown-menu label="Code slot" dynamic-align="" class="flex">
            <paper-listbox slot="dropdown-content" selected="{{selectedUserCode}}">
              <template is="dom-repeat" items="[[userCodes]]" as="state">
                <paper-item>[[computeSelectCaptionUserCodes(state)]]</paper-item>
              </template>
            </paper-listbox>
          </paper-dropdown-menu>
          </div>

          <template is="dom-if" if="[[isUserCodeSelected(selectedUserCode)]]">
            <div class="card-actions">
              <paper-input label="User code" type="text" allowed-pattern="[0-9,a-f,x,\\\\]" maxlength="{{userCodeMaxLen}}" minlength="16" value="{{selectedUserCodeValue}}">
              </paper-input>
              <pre>Ascii: [[computedCodeOutput]]</pre>
            </div>
            <div class="card-actions">
              <ha-call-service-button hass="[[hass]]" domain="lock" service="set_usercode" service-data="[[computeUserCodeServiceData(selectedUserCodeValue, &quot;Add&quot;)]]">Set Usercode</ha-call-service-button>
              <ha-call-service-button hass="[[hass]]" domain="lock" service="clear_usercode" service-data="[[computeUserCodeServiceData(selectedUserCode, &quot;Delete&quot;)]]">Delete Usercode</ha-call-service-button>
            </div>
          </template>
        </paper-card>
      </div>
`}static get properties(){return{hass:{type:Object},nodes:{type:Array},selectedNode:{type:Number},userCodes:{type:Object},userCodeMaxLen:{type:Number,value:4},selectedUserCode:{type:Number,value:-1,observer:"selectedUserCodeChanged"},selectedUserCodeValue:{type:String},computedCodeOutput:{type:String,value:""}}}ready(){super.ready(),this.addEventListener("hass-service-called",e=>this.serviceCalled(e))}serviceCalled(e){if(e.detail.success){var t=this;setTimeout(function(){t.refreshUserCodes(t.selectedNode)},5e3)}}isUserCodeSelected(e){return-1!==e}computeSelectCaptionUserCodes(e){return e.key+": "+e.value.label}selectedUserCodeChanged(e){if(-1!==this.selectedUserCode&&-1!==e){var t=this.userCodes[e].value.code;this.userCodeMaxLen=4*this.userCodes[e].value.length,this.selectedUserCodeValue=this.a2hex(t),this.computedCodeOutput="["+this.hex2a(this.selectedUserCodeValue)+"]"}}computeUserCodeServiceData(e,t){if(-1===this.selectedNode||!e)return-1;var i=null,a=null;return"Add"===t&&(a=this.hex2a(e),this.computedCodeOutput="["+a+"]",i={node_id:this.nodes[this.selectedNode].attributes.node_id,code_slot:this.selectedUserCode,usercode:a}),"Delete"===t&&(i={node_id:this.nodes[this.selectedNode].attributes.node_id,code_slot:this.selectedUserCode}),i}refreshUserCodes(e){this.selectedUserCodeValue="";var t=[];this.hass.callApi("GET","zwave/usercodes/"+this.nodes[e].attributes.node_id).then(function(e){Object.keys(e).forEach(function(i){t.push({key:i,value:e[i]})}),this.userCodes=t,this.selectedUserCodeChanged(this.selectedUserCode)}.bind(this))}a2hex(e){for(var t=[],i="",a=0,s=e.length;a<s;a++){var o=Number(e.charCodeAt(a)).toString(16);i="0"===o?"00":o,t.push("\\x"+i)}return t.join("")}hex2a(e){for(var t=e.toString().replace(/\\x/g,""),i="",a=0;a<t.length;a+=2)i+=String.fromCharCode(parseInt(t.substr(a,2),16));return i}});customElements.define("zwave-values",class extends s.a{static get template(){return a["a"]`
    <style include="iron-flex ha-style">
      .content {
        margin-top: 24px;
      }

      paper-card {
        display: block;
        margin: 0 auto;
        max-width: 600px;
      }

      .device-picker {
        @apply --layout-horizontal;
        @apply --layout-center-center;
        padding-left: 24px;
        padding-right: 24px;
        padding-bottom: 24px;
        }

      .help-text {
        padding-left: 24px;
        padding-right: 24px;
      }
    </style>
    <div class="content">
      <paper-card heading="Node Values">
        <div class="device-picker">
        <paper-dropdown-menu label="Value" dynamic-align="" class="flex">
          <paper-listbox slot="dropdown-content" selected="{{selectedValue}}">
             <template is="dom-repeat" items="[[values]]" as="item">
              <paper-item>[[computeSelectCaption(item)]]</paper-item>
            </template>
          </paper-listbox>
        </paper-dropdown-menu>
        </div>
        <template is="dom-if" if="[[!computeIsValueSelected(selectedValue)]]">
          <div class="card-actions">
            <paper-input float-label="Value Name" type="text" value="{{newValueNameInput}}" placeholder="[[computeGetValueName(selectedValue)]]">
            </paper-input>
            <ha-call-service-button hass="[[hass]]" domain="zwave" service="rename_value" service-data="[[computeValueNameServiceData(newValueNameInput)]]">Rename Value</ha-call-service-button>

          </div>
        </template>
      </paper-card>
    </div>
`}static get properties(){return{hass:{type:Object},nodes:{type:Array},values:{type:Array},selectedNode:{type:Number},selectedValue:{type:Number,value:-1,observer:"selectedValueChanged"}}}ready(){super.ready(),this.addEventListener("hass-service-called",e=>this.serviceCalled(e))}serviceCalled(e){var t=this;e.detail.success&&setTimeout(function(){t.refreshValues(t.selectedNode)},5e3)}computeSelectCaption(e){return e.value.label+" (Instance: "+e.value.instance+", Index: "+e.value.index+")"}computeGetValueName(e){return this.values[e].value.label}computeIsValueSelected(e){return!this.nodes||-1===this.selectedNode||-1===e}refreshValues(e){var t=[];this.hass.callApi("GET","zwave/values/"+this.nodes[e].attributes.node_id).then(function(e){Object.keys(e).forEach(function(i){t.push({key:i,value:e[i]})}),this.values=t,this.selectedValueChanged(this.selectedValue)}.bind(this))}computeValueNameServiceData(e){return-1===!this.selectedNode||-1===this.selectedValue?-1:{node_id:this.nodes[this.selectedNode].attributes.node_id,value_id:this.values[this.selectedValue].key,name:e}}selectedValueChanged(e){if(-1!==!this.selectedNode&&-1!==this.selectedValue){var t=this;this.hass.callApi("GET","config/zwave/device_config/"+this.values[e].value.entity_id).then(function(i){t.entityIgnored=i.ignored||!1,t.entityPollingIntensity=t.values[e].value.poll_intensity})}}});customElements.define("ha-config-zwave",class extends(window.hassMixins.LocalizeMixin(s.a)){static get template(){return a["a"]`
    <style include="iron-flex ha-style ha-form-style">
      .content {
        margin-top: 24px;
      }

      .node-info {
        margin-left: 16px;
      }

      .help-text {
        padding-left: 24px;
        padding-right: 24px;
      }

      paper-card {
        display: block;
        margin: 0 auto;
        max-width: 600px;
      }

      .device-picker {
        @apply --layout-horizontal;
        @apply --layout-center-center;
        padding-left: 24px;
        padding-right: 24px;
        padding-bottom: 24px;
      }

      ha-service-description {
        display: block;
        color: grey;
      }

      [hidden] {
        display: none;
      }

      .toggle-help-icon {
        position: absolute;
        top: 6px;
        right: 0;
        color: var(--primary-color);
      }
    </style>
    <ha-app-layout has-scrolling-region="">
      <app-header slot="header" fixed="">
        <app-toolbar>
          <paper-icon-button icon="mdi:arrow-left" on-click="_backTapped"></paper-icon-button>
          <div main-title="">[[localize('ui.panel.config.zwave.caption')]]</div>
        </app-toolbar>
      </app-header>

      <zwave-network id="zwave-network" is-wide="[[isWide]]" hass="[[hass]]"></zwave-network>

      <!--Node card-->
      <ha-config-section is-wide="[[isWide]]">
        <div style="position: relative" slot="header">
          <span>Z-Wave Node Management</span>
          <paper-icon-button class="toggle-help-icon" on-click="toggleHelp" icon="mdi:help-circle"></paper-icon-button>

        </div>
        <span slot="introduction">
          Run Z-Wave commands that affect a single node. Pick a node to see a list of available commands.
        </span>

        <paper-card class="content">
          <div class="device-picker">
            <paper-dropdown-menu dynamic-align="" label="Nodes" class="flex">
              <paper-listbox slot="dropdown-content" selected="{{selectedNode}}">
                <template is="dom-repeat" items="[[nodes]]" as="state">
                  <paper-item>[[computeSelectCaption(state)]]</paper-item>
                </template>
              </paper-listbox>
            </paper-dropdown-menu>
          </div>
            <template is="dom-if" if="[[!computeIsNodeSelected(selectedNode)]]">
              <template is="dom-if" if="[[showHelp]]">
                <div style="color: grey; padding: 12px">Select node to view per-node options</div>
              </template>
            </template>

          <template is="dom-if" if="[[computeIsNodeSelected(selectedNode)]]">
          <div class="card-actions">
            <ha-call-service-button hass="[[hass]]" domain="zwave" service="refresh_node" service-data="[[computeNodeServiceData(selectedNode)]]">Refresh Node</ha-call-service-button>
            <ha-service-description hass="[[hass]]" domain="zwave" service="refresh_node" hidden\$="[[!showHelp]]"></ha-service-description>

            <ha-call-service-button hass="[[hass]]" domain="zwave" service="remove_failed_node" service-data="[[computeNodeServiceData(selectedNode)]]">Remove Failed Node</ha-call-service-button>
            <ha-service-description hass="[[hass]]" domain="zwave" service="remove_failed_node" hidden\$="[[!showHelp]]"></ha-service-description>

            <ha-call-service-button hass="[[hass]]" domain="zwave" service="replace_failed_node" service-data="[[computeNodeServiceData(selectedNode)]]">Replace Failed Node</ha-call-service-button>
            <ha-service-description hass="[[hass]]" domain="zwave" service="replace_failed_node" hidden\$="[[!showHelp]]"></ha-service-description>

            <ha-call-service-button hass="[[hass]]" domain="zwave" service="print_node" service-data="[[computeNodeServiceData(selectedNode)]]">Print Node</ha-call-service-button>
            <ha-service-description hass="[[hass]]" domain="zwave" service="print_node" hidden\$="[[!showHelp]]"></ha-service-description>

            <ha-call-service-button hass="[[hass]]" domain="zwave" service="heal_node" service-data="[[computeHealNodeServiceData(selectedNode)]]">Heal Node</ha-call-service-button>
            <ha-service-description hass="[[hass]]" domain="zwave" service="heal_node" hidden\$="[[!showHelp]]"></ha-service-description>

            <ha-call-service-button hass="[[hass]]" domain="zwave" service="test_node" service-data="[[computeNodeServiceData(selectedNode)]]">Test Node</ha-call-service-button>
            <ha-service-description hass="[[hass]]" domain="zwave" service="test_node" hidden\$="[[!showHelp]]"></ha-service-description>
          </div>
          <div class="card-actions">
            <paper-input float-label="New node name" type="text" value="{{newNodeNameInput}}" placeholder="[[computeGetNodeName(selectedNode)]]">
            </paper-input>
            <ha-call-service-button hass="[[hass]]" domain="zwave" service="rename_node" service-data="[[computeNodeNameServiceData(newNodeNameInput)]]">Rename Node</ha-call-service-button>
            <ha-service-description hass="[[hass]]" domain="zwave" service="rename_node" hidden\$="[[!showHelp]]"></ha-service-description>
           </div>

           <div class="device-picker">
            <paper-dropdown-menu label="Entities of this node" dynamic-align="" class="flex">
              <paper-listbox slot="dropdown-content" selected="{{selectedEntity}}">
                <template is="dom-repeat" items="[[entities]]" as="state">
                  <paper-item>[[computeSelectCaptionEnt(state)]]</paper-item>
                </template>
              </paper-listbox>
            </paper-dropdown-menu>
           </div>
           <template is="dom-if" if="[[!computeIsEntitySelected(selectedEntity)]]">
           <div class="card-actions">
             <ha-call-service-button hass="[[hass]]" domain="zwave" service="refresh_entity" service-data="[[computeRefreshEntityServiceData(selectedEntity)]]">Refresh Entity</ha-call-service-button>
             <ha-service-description hass="[[hass]]" domain="zwave" service="refresh_entity" hidden\$="[[!showHelp]]"></ha-service-description>
           </div>
           <div class="form-group">
             <paper-checkbox checked="{{entityIgnored}}" class="form-control">
             Exclude this entity from Home Assistant
             </paper-checkbox>
             <paper-input disabled="{{entityIgnored}}" label="Polling intensity" type="number" min="0" value="{{entityPollingIntensity}}">
             </paper-input>
           </div>
           <div class="card-actions">
             <ha-call-service-button hass="[[hass]]" domain="zwave" service="set_poll_intensity" service-data="[[computePollIntensityServiceData(entityPollingIntensity)]]">Save</ha-call-service-button>
           </div>
           <div class="content">
             <div class="card-actions">
               <paper-button toggles="" raised="" noink="" active="{{entityInfoActive}}">Entity Attributes</paper-button>
             </div>
             <template is="dom-if" if="{{entityInfoActive}}">
               <template is="dom-repeat" items="[[selectedEntityAttrs]]" as="state">
                 <div class="node-info">
                   <span>[[state]]</span>
                 </div>
               </template>
             </template>
           </div>

           </template>
          </template>
        </paper-card>

        <template is="dom-if" if="[[computeIsNodeSelected(selectedNode)]]">
          <!--Node info card-->
          <zwave-node-information id="zwave-node-information" nodes="[[nodes]]" selected-node="[[selectedNode]]"></zwave-node-information>

          <!--Value card-->
          <zwave-values hass="[[hass]]" nodes="[[nodes]]" selected-node="[[selectedNode]]" values="[[values]]"></zwave-values>

          <!--Group card-->
          <zwave-groups hass="[[hass]]" nodes="[[nodes]]" selected-node="[[selectedNode]]" groups="[[groups]]"></zwave-groups>

          <!--Config card-->
          <zwave-node-config hass="[[hass]]" nodes="[[nodes]]" selected-node="[[selectedNode]]" config="[[config]]"></zwave-node-config>
        </template>

        <!--User Codes-->
        <template is="dom-if" if="{{hasNodeUserCodes}}">
          <zwave-usercodes id="zwave-usercodes" hass="[[hass]]" nodes="[[nodes]]" user-codes="[[userCodes]]" selected-node="[[selectedNode]]"></zwave-usercodes>
      </template>
      </ha-config-section>



      <!--Ozw log-->
      <ozw-log is-wide="[[isWide]]" hass="[[hass]]"></ozw-log>

    </ha-app-layout>
`}static get properties(){return{hass:Object,isWide:Boolean,nodes:{type:Array,computed:"computeNodes(hass)"},selectedNode:{type:Number,value:-1,observer:"selectedNodeChanged"},config:{type:Array,value:function(){return[]}},entities:{type:Array,computed:"computeEntities(selectedNode)"},entityInfoActive:{type:Boolean},selectedEntity:{type:Number,value:-1,observer:"selectedEntityChanged"},selectedEntityAttrs:{type:Array,computed:"computeSelectedEntityAttrs(selectedEntity)"},values:{type:Array},groups:{type:Array},newNodeNameInput:{type:String},userCodes:{type:Array,value:function(){return[]}},hasNodeUserCodes:{type:Boolean,value:!1},showHelp:{type:Boolean,value:!1},entityIgnored:{type:Boolean},entityPollingIntensity:{type:Number,value:0}}}ready(){super.ready(),this.addEventListener("hass-service-called",e=>this.serviceCalled(e))}serviceCalled(e){e.detail.success&&"set_poll_intensity"===e.detail.service&&this.saveEntity()}computeNodes(e){return Object.keys(e.states).map(function(t){return e.states[t]}).filter(function(e){return e.entity_id.match("zwave[.]")}).sort(ae)}computeEntities(e){if(!this.nodes||-1===e)return-1;var t=this.hass,i=this.nodes[this.selectedNode].attributes.node_id;return Object.keys(t.states).map(function(e){return t.states[e]}).filter(function(e){return void 0!==e.attributes.node_id&&!e.attributes.hidden&&"node_id"in e.attributes&&e.attributes.node_id===i&&!e.entity_id.match("zwave[.]")}).sort(ae)}selectedNodeChanged(e){this.newNodeNameInput="",-1!==e&&(this.selectedConfigParameter=-1,this.selectedConfigParameterValue=-1,this.selectedGroup=-1,this.hass.callApi("GET","zwave/config/"+this.nodes[e].attributes.node_id).then(e=>{this.config=this._objToArray(e)}),this.hass.callApi("GET","zwave/values/"+this.nodes[e].attributes.node_id).then(e=>{this.values=this._objToArray(e)}),this.hass.callApi("GET","zwave/groups/"+this.nodes[e].attributes.node_id).then(e=>{this.groups=this._objToArray(e)}),this.hasNodeUserCodes=!1,this.notifyPath("hasNodeUserCodes"),this.hass.callApi("GET","zwave/usercodes/"+this.nodes[e].attributes.node_id).then(e=>{this.userCodes=this._objToArray(e),this.hasNodeUserCodes=this.userCodes.length>0,this.notifyPath("hasNodeUserCodes")}))}selectedEntityChanged(e){if(-1!==e){var t=this;t.hass.callApi("GET","zwave/values/"+t.nodes[t.selectedNode].attributes.node_id).then(e=>{t.values=t._objToArray(e)});var i=t.entities[e].attributes.value_id,a=t.values.find(function(e){return e.key===i}),s=t.values.indexOf(a);t.hass.callApi("GET","config/zwave/device_config/"+i).then(function(e){t.entityIgnored=e.ignored||!1,t.entityPollingIntensity=t.values[s].value.poll_intensity})}}computeSelectedEntityAttrs(e){if(-1===e)return"No entity selected";var t=this.entities[e].attributes,i=[];return Object.keys(t).forEach(function(e){i.push(e+": "+t[e])}),i.sort()}computeSelectCaption(e){return Object(Z.a)(e)+" (Node:"+e.attributes.node_id+" "+e.attributes.query_stage+")"}computeSelectCaptionEnt(e){return Object(f.a)(e)+"."+Object(Z.a)(e)}computeIsNodeSelected(){return this.nodes&&-1!==this.selectedNode}computeIsEntitySelected(e){return-1===e}computeNodeServiceData(e){return{node_id:this.nodes[e].attributes.node_id}}computeHealNodeServiceData(e){return{node_id:this.nodes[e].attributes.node_id,return_routes:!0}}computeGetNodeName(e){return-1!==this.selectedNode&&this.nodes[e].entity_id?this.nodes[e].attributes.node_name:-1}computeNodeNameServiceData(e){return{node_id:this.nodes[this.selectedNode].attributes.node_id,name:e}}computeRefreshEntityServiceData(e){return-1===e?-1:{entity_id:this.entities[e].entity_id}}computePollIntensityServiceData(e){return-1===!this.selectedNode||-1===this.selectedEntity?-1:{node_id:this.nodes[this.selectedNode].attributes.node_id,value_id:this.entities[this.selectedEntity].attributes.value_id,poll_intensity:parseInt(e)}}saveEntity(){var e={ignored:this.entityIgnored,polling_intensity:parseInt(this.entityPollingIntensity)};return this.hass.callApi("POST","config/zwave/device_config/"+this.entities[this.selectedEntity].entity_id,e)}toggleHelp(){this.showHelp=!this.showHelp}_objToArray(e){var t=[];return Object.keys(e).forEach(function(i){t.push({key:i,value:e[i]})}),t}_backTapped(){history.back()}});customElements.define("ha-panel-config",class extends(window.hassMixins.NavigateMixin(s.a)){static get template(){return a["a"]`
    <app-route
      route='[[route]]'
      pattern='/:page'
      data="{{_routeData}}"
    ></app-route>

    <iron-media-query query="(min-width: 1040px)" query-matches="{{wide}}">
    </iron-media-query>
    <iron-media-query query="(min-width: 1296px)" query-matches="{{wideSidebar}}">
    </iron-media-query>

    <template is="dom-if" if='[[_equals(_routeData.page, "core")]]' restamp>
      <ha-config-core
        page-name='core'
        hass='[[hass]]'
        is-wide='[[isWide]]'
      ></ha-config-core>
    </template>

    <template is="dom-if" if='[[_equals(_routeData.page, "cloud")]]' restamp>
      <ha-config-cloud
        page-name='cloud'
        route='[[route]]'
        hass='[[hass]]'
        is-wide='[[isWide]]'
        account='[[account]]'
      ></ha-config-cloud>
    </template>

    <template is="dom-if" if='[[_equals(_routeData.page, "dashboard")]]'>
      <ha-config-dashboard
        page-name='dashboard'
        hass='[[hass]]'
        is-wide='[[isWide]]'
        account='[[account]]'
        narrow='[[narrow]]'
        show-menu='[[showMenu]]'
      ></ha-config-dashboard>
    </template>

    <template is="dom-if" if='[[_equals(_routeData.page, "automation")]]' restamp>
      <ha-config-automation
        page-name='automation'
        route='[[route]]'
        hass='[[hass]]'
        is-wide='[[isWide]]'
      ></ha-config-automation>
    </template>

    <template is="dom-if" if='[[_equals(_routeData.page, "script")]]' restamp>
      <ha-config-script
        page-name='script'
        route='[[route]]'
        hass='[[hass]]'
        is-wide='[[isWide]]'
      ></ha-config-script>
    </template>

    <template is="dom-if" if='[[_equals(_routeData.page, "zwave")]]' restamp>
      <ha-config-zwave
        page-name='zwave'
        hass='[[hass]]'
        is-wide='[[isWide]]'
      ></ha-config-zwave>
    </template>

    <template is="dom-if" if='[[_equals(_routeData.page, "customize")]]' restamp>
      <ha-config-customize
        page-name='customize'
        hass='[[hass]]'
        is-wide='[[isWide]]'
      ></ha-config-customize>
    </template>

    <template is="dom-if" if='[[_equals(_routeData.page, "integrations")]]' restamp>
      <ha-config-entries
        page-name='integrations'
        hass='[[hass]]'
        is-wide='[[isWide]]'
      ></ha-config-entries>
    </template>
    `}static get properties(){return{hass:Object,narrow:Boolean,showMenu:Boolean,account:Object,route:{type:Object,observer:"_routeChanged"},_routeData:Object,wide:Boolean,wideSidebar:Boolean,isWide:{type:Boolean,computed:"computeIsWide(showMenu, wideSidebar, wide)"}}}ready(){super.ready(),Object(ee.a)(this.hass,"cloud")&&this.hass.callApi("get","cloud/account").then(e=>{this.account=e}),this.addEventListener("ha-account-refreshed",e=>{this.account=e.detail.account})}computeIsWide(e,t,i){return e?t:i}_routeChanged(e){""===e.path&&"/config"===e.prefix&&this.navigate("/config/dashboard",!0)}_equals(e,t){return e===t}})}}]);