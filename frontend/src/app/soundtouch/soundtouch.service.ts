import { ApiService } from './../api/api.service';
import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import { Observable } from 'rxjs/Rx';

@Injectable()
export class SoundtouchService {

  now_playing : string = ':8090/now_playing';

  constructor(private http: Http,private apiService : ApiService) {}

  getNowPlaying(){
    let url = this.apiService.getUrl('soundtouch') + this.now_playing;
    console.log(url);
    return this.http.get(url)
      .map(res => {  
        let parser=new DOMParser();
        let xmlDoc = parser.parseFromString(res.text(),"text/xml");
        if(xmlDoc.getElementsByTagName("track")[0]){
          if(xmlDoc.getElementsByTagName("ContentItem")[0].getAttribute("source") == "BLUETOOTH"){
            return {
              playing : true,
              source : xmlDoc.getElementsByTagName("ContentItem")[0].getAttribute("source"),
              track : xmlDoc.getElementsByTagName("stationName")[0].childNodes[0].nodeValue
            }
          }else{
            return {
              playing : true,
              source : xmlDoc.getElementsByTagName("ContentItem")[0].getAttribute("source"),
              artist : xmlDoc.getElementsByTagName("artist")[0].childNodes[0].nodeValue,
              track : xmlDoc.getElementsByTagName("track")[0].childNodes[0].nodeValue,
              album : xmlDoc.getElementsByTagName("album")[0].childNodes[0].nodeValue,
              art : xmlDoc.getElementsByTagName("art")[0].childNodes[0].nodeValue
            }
          }
        }
        else if(xmlDoc.getElementsByTagName("ContentItem")[0].getAttribute("source") == 'STANDBY'){
          return {
            playing : false,
            message : "STANDBY"
          }
        }else{
          return {
            playing : false,
            message : "STATUT OFF"
          } 
        }
      });
  }


}
