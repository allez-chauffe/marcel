import {Injectable} from '@angular/core';
import { Http } from '@angular/http';

@Injectable()
export class VlilleService {

  //view-source:http://vlille.fr/stations/xml-stations.aspx
  //http://vlille.fr/stations/xml-station.aspx?borne=10 //station rihour
  //http://vlille.fr/stations/xml-station.aspx?borne=36 //station cormontaigne

  public api_url : String = "http://api.aurelien-loyer.fr/vlille/station.php?key=magdalena";

  constructor(public http : Http) {
  }

  getBorneData(borne){
    return this.http
      .get(this.api_url+'&format=jsonp&json_callback=JSON_CALLBACK&borne='+borne.id)
      .map(function(res){
        const response = res.json();
        response.name = borne.name;
        return response;
      });
  }
  
}
