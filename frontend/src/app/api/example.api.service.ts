//Remove exemple. in the filename :)

import {Injectable} from '@angular/core';

@Injectable()
export class ApiService {

  public apis : any;

  constructor() {

    this.apis = {
      soundtouch : {
        url : ''
      },
      maps : {
        url : 'http://localhost:8090/api/v1/weather/forecast/2',
        key : 'YOUR_MAPS_API_KEY'
      },
      weather : {
        key : 'YOUR_FREE_OPENWEATHER_API_KEY'
      },
      sparrow: {
        url : 'http://10.0.10.245/arduino/getRoomStat'
      },
      smaug: {
        url : 'http://10.0.10.40/arduino/getRoomStat'
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
