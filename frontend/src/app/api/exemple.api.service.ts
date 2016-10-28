//Remove exemple. in the filename :)

import {Injectable} from '@angular/core';

@Injectable()
export class ApiService {

  public apis : any;

  constructor() {

    this.apis = {
      maps : {
        key : 'YOUR_MAPS_API_KEY'
      },
      weather : {
        key : 'YOUR_FREE_OPENWEATHER_API_KEY'
      },
      xee : {
        url : 'http://RpiServerIp/domoticz_scripts/xee-car-data-to-domoticz-php/xee.php?data='
        // Projet Github
        // @t3kstiil3
        // https://github.com/T3kstiil3/Domoticz_Scripts/tree/master/xee-car-data-to-domoticz-php
      }
    };
  }

  getKeys(){
    return this.apis;
  }

  getApi(name){
    return this.apis[name];
  }

  getKey(name){
    return this.apis[name].key;
  }

  getUrl(name){
    return this.apis[name].url;
  }

}
