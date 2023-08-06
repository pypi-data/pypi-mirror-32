(window.webpackJsonp=window.webpackJsonp||[]).push([[9],{273:function(t,a,e){"use strict";e.r(a),e(117),e(50);var i=e(0),r=e(3),s=e(214),n=e.n(s),o=(e(131),e(145),e(11));customElements.define("ha-entity-marker",class extends(Object(o.a)(r.a)){static get template(){return i["a"]`
    <style include="iron-positioning"></style>
    <style>
    .marker {
      vertical-align: top;
      position: relative;
      display: block;
      margin: 0 auto;
      width: 2.5em;
      text-align: center;
      height: 2.5em;
      line-height: 2.5em;
      font-size: 1.5em;
      border-radius: 50%;
      border: 0.1em solid var(--ha-marker-color, var(--default-primary-color));
      color: rgb(76, 76, 76);
      background-color: white;
    }
    iron-image {
      border-radius: 50%;
    }
    </style>

    <div class="marker">
      <template is="dom-if" if="[[entityName]]">[[entityName]]</template>
      <template is="dom-if" if="[[entityPicture]]">
        <iron-image sizing="cover" class="fit" src="[[entityPicture]]"></iron-image>
      </template>
    </div>
`}static get properties(){return{hass:{type:Object},entityId:{type:String,value:""},entityName:{type:String,value:null},entityPicture:{type:String,value:null}}}ready(){super.ready(),this.addEventListener("click",t=>this.badgeTap(t))}badgeTap(t){t.stopPropagation(),this.entityId&&this.fire("hass-more-info",{entityId:this.entityId})}});var m=e(14),c=e(15),h=e(19);n.a.Icon.Default.imagePath="/static/images/leaflet",customElements.define("ha-panel-map",class extends(Object(h.a)(r.a)){static get template(){return i["a"]`
    <style include="ha-style">
      #map {
        height: calc(100% - 64px);
        width: 100%;
        z-index: 0;
      }
    </style>

    <app-toolbar>
      <ha-menu-button narrow='[[narrow]]' show-menu='[[showMenu]]'></ha-menu-button>
      <div main-title>[[localize('panel.map')]]</div>
    </app-toolbar>

    <div id='map'></div>
    `}static get properties(){return{hass:{type:Object,observer:"drawEntities"},narrow:{type:Boolean},showMenu:{type:Boolean,value:!1}}}connectedCallback(){super.connectedCallback();var t=n.a.tileLayer("http://webrd02.is.autonavi.com/appmaptile?lang=zh_cn&size=1&scale=1&style=8&x={x}&y={y}&z={z}",{attribution:'© <a href="http://map.amap.com/doc/serviceitem.html">高德地图</a> contributors, © <a href="http://lbs.amap.com/">高德软件</a>',maxZoom:18}),a=n.a.tileLayer("http://webst04.is.autonavi.com/appmaptile?style=6&x={x}&y={y}&z={z}",{attribution:'© <a href="http://map.amap.com/doc/serviceitem.html">高德地图</a> contributors, © <a href="http://lbs.amap.com/">高德软件</a>',maxZoom:18}),e=n.a.tileLayer("http://tm.amap.com/trafficengine/mapabc/traffictile?v=1.0&;t=1&zoom={z}&x={x}&y={y}",{maxZoom:17,zoomReverse:!0}),i={"街道":t,"卫星":a},r={"路况":e},s=this._map=n.a.map(this.$.map,{layers:[t,e]}),o=document.createElement("link");o.setAttribute("href","/static/images/leaflet/leaflet.css"),o.setAttribute("rel","stylesheet"),this.$.map.parentNode.appendChild(o),s.setView([51.505,-.09],13),n.a.control.layers(i,r).addTo(s),this.drawEntities(this.hass),setTimeout(()=>{s.invalidateSize(),this.fitMap()},1)}fitMap(){var t;0===this._mapItems.length?this._map.setView(new n.a.LatLng(this.hass.config.core.latitude,this.hass.config.core.longitude),14):(t=new n.a.latLngBounds(this._mapItems.map(t=>t.getLatLng())),this._map.fitBounds(t.pad(.5)))}drawEntities(t){var a=function(){};a.prototype.x_PI=52.35987755982988,a.prototype.a=6378245,a.prototype.ee=.006693421622965943,a.prototype.transform=function(t,a,e){if(this.outOfChina(t,a))return[t,a];if(dLat=this.transformLat(a-105,t-35),dLon=this.transformLon(a-105,t-35),radLat=t/180*Math.PI,magic=Math.sin(radLat),magic=1-this.ee*magic*magic,sqrtMagic=Math.sqrt(magic),dLat=180*dLat/(this.a*(1-this.ee)/(magic*sqrtMagic)*Math.PI),dLon=180*dLon/(this.a/sqrtMagic*Math.cos(radLat)*Math.PI),mgLat=t+dLat,mgLon=a+dLon,e)return[mgLat,mgLon];var i=Math.sqrt(mgLon*mgLon+mgLat*mgLat)+2e-5*Math.sin(mgLat*this.x_PI),r=Math.atan2(mgLat,mgLon)+3e-6*Math.cos(mgLon*this.x_PI),s=i*Math.cos(r)+.0065;return[i*Math.sin(r)+.006,s]},a.prototype.outOfChina=function(t,a){return a<72.004||a>137.8347||t<.8293||t>55.8271},a.prototype.transformLat=function(t,a){var e=2*t-100+3*a+.2*a*a+.1*t*a+.2*Math.sqrt(Math.abs(t));return e+=2*(20*Math.sin(6*t*Math.PI)+20*Math.sin(2*t*Math.PI))/3,(e+=2*(20*Math.sin(a*Math.PI)+40*Math.sin(a/3*Math.PI))/3)+2*(160*Math.sin(a/12*Math.PI)+320*Math.sin(a*Math.PI/30))/3},a.prototype.transformLon=function(t,a){var e=300+t+2*a+.1*t*t+.1*t*a+.1*Math.sqrt(Math.abs(t));return e+=2*(20*Math.sin(6*t*Math.PI)+20*Math.sin(2*t*Math.PI))/3,(e+=2*(20*Math.sin(t*Math.PI)+40*Math.sin(t/3*Math.PI))/3)+2*(150*Math.sin(t/12*Math.PI)+300*Math.sin(t/30*Math.PI))/3};var e=this._map;if(e){this._mapItems&&this._mapItems.forEach(function(t){t.remove()});var i=this._mapItems=[];Object.keys(t.states).forEach(function(r){var s=t.states[r],o=Object(c.a)(s);if(!(s.attributes.hidden&&"zone"!==Object(m.a)(s)||"home"===s.state)&&"latitude"in s.attributes&&"longitude"in s.attributes){var h;if("zone"===Object(m.a)(s)){if(s.attributes.passive)return;var l;return l=s.attributes.icon?"<iron-icon icon='"+s.attributes.icon+"'></iron-icon>":o,h=n.a.divIcon({html:l,iconSize:[24,24],className:""}),i.push(n.a.marker((new a).transform(s.attributes.latitude,s.attributes.longitude,!1),{icon:h,interactive:!1,title:o}).addTo(e)),void i.push(n.a.circle((new a).transform(s.attributes.latitude,s.attributes.longitude,!1),{interactive:!1,color:"#FF9800",radius:s.attributes.radius}).addTo(e))}var p=s.attributes.entity_picture||"",u=o.split(" ").map(function(t){return t.substr(0,1)}).join("");h=n.a.divIcon({html:"<ha-entity-marker entity-id='"+s.entity_id+"' entity-name='"+u+"' entity-picture='"+p+"'></ha-entity-marker>",iconSize:[45,45],className:""}),i.push(n.a.marker((new a).transform(s.attributes.latitude,s.attributes.longitude,!1),{icon:h,title:Object(c.a)(s)}).addTo(e)),s.attributes.gps_accuracy&&i.push(n.a.circle((new a).transform(s.attributes.latitude,s.attributes.longitude,!1),{interactive:!1,color:"#0288D1",radius:s.attributes.gps_accuracy}).addTo(e))}})}}})}}]);
//# sourceMappingURL=ab1bde70e912f9275b6a.chunk.js.map