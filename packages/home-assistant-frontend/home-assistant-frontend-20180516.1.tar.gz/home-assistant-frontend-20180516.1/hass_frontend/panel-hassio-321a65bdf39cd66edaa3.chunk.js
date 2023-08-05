(window.webpackJsonp=window.webpackJsonp||[]).push([[14],{398:function(e,t,s){"use strict";s.r(t);var a=s(0),i=s(3);s(14);customElements.define("ha-panel-hassio",class extends(window.hassMixins.NavigateMixin(window.hassMixins.EventsMixin(i.a))){static get template(){return a["a"]`
    <style>
      iframe {
        border: 0;
        width: 100%;
        height: 100%;
        display: block;
      }
    </style>
    <iframe
      id='iframe'
      src="[[iframeUrl]]"
    ></iframe>
    `}static get properties(){return{hass:Object,narrow:Boolean,showMenu:Boolean,route:Object,iframeUrl:{type:String,value:"/api/hassio/app-es5/index.html"}}}static get observers(){return["_dataChanged(hass, narrow, showMenu, route)"]}ready(){super.ready(),window.hassioPanel=this}_dataChanged(e,t,s,a){this._updateProperties({hass:e,narrow:t,showMenu:s,route:a})}_updateProperties(e){const t=this.$.iframe.contentWindow&&this.$.iframe.contentWindow.setProperties;if(!t){const t=!this._dataToSet;return this._dataToSet=e,void(t&&setTimeout(()=>{const e=this._dataToSet;this._dataToSet=null,this._updateProperties(e)},100))}t(e)}})}}]);