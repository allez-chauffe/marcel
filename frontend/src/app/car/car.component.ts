import {Component, OnInit,ElementRef} from "@angular/core";
import {ViewEncapsulation} from '@angular/core';
import {NgStyle} from '@angular/common';
import {CarService} from './car.service';
import {Observable} from 'rxjs/Rx';
import {ApiService} from './../api/api.service';

@Component({
  selector: 'car',
  encapsulation: ViewEncapsulation.None,
  templateUrl: './car.component.html',
  styleUrls: ['./car.component.scss']
})

export class CarComponent implements OnInit{

  public lat:number;
  public lng:number;
  public latCar:number;
  public lngCar:number;
  public zoom:number;
  public map:any;
  public styles:any;
  public loop:any;
  public state:string;
  public carIcon:string;
  public positionInfoMarker:string;
  public positionInfo:string;
  public fuelLevel : number;
  public fuelRange : number;
  public trips : any[] = [];
  public apikey : String;
  public interval : number;
  public numberOfTrips : number;
  public engineSpeed : number = 0;
  public lastUpdateDate : any;
  public positionClass : string = "";

  constructor(public elementRef:ElementRef,public carService : CarService, public apiService : ApiService) {
    this.interval = 30000; // 30 secondes
    this.lat = 50.631941;
    this.lng = 3.057928;
    this.latCar = 50.626301;
    this.lngCar = 3.032630;
    this.zoom = 12;
    this.map = elementRef.nativeElement.querySelector("#gmap");
    this.carIcon = "assets/ds3_40px.png";
    this.styles = [{"featureType":"all","elementType":"labels","stylers":[{"visibility":"on"}]},{"featureType":"all","elementType":"labels.text.fill","stylers":[{"saturation":36},{"color":"#000000"},{"lightness":40}]},{"featureType":"all","elementType":"labels.text.stroke","stylers":[{"visibility":"on"},{"color":"#000000"},{"lightness":16}]},{"featureType":"all","elementType":"labels.icon","stylers":[{"visibility":"off"}]},{"featureType":"administrative","elementType":"geometry.fill","stylers":[{"color":"#000000"},{"lightness":20}]},{"featureType":"administrative","elementType":"geometry.stroke","stylers":[{"color":"#000000"},{"lightness":17},{"weight":1.2}]},{"featureType":"administrative.country","elementType":"labels.text.fill","stylers":[{"color":"#ed5929"}]},{"featureType":"administrative.locality","elementType":"labels.text.fill","stylers":[{"color":"#c4c4c4"}]},{"featureType":"administrative.neighborhood","elementType":"labels.text.fill","stylers":[{"color":"#ed5929"}]},{"featureType":"landscape","elementType":"geometry","stylers":[{"color":"#000000"},{"lightness":20}]},{"featureType":"poi","elementType":"geometry","stylers":[{"color":"#000000"},{"lightness":21},{"visibility":"on"}]},{"featureType":"poi.business","elementType":"geometry","stylers":[{"visibility":"on"}]},{"featureType":"road.highway","elementType":"geometry.fill","stylers":[{"color":"#ed5929"},{"lightness":"0"}]},{"featureType":"road.highway","elementType":"geometry.stroke","stylers":[{"visibility":"off"}]},{"featureType":"road.highway","elementType":"labels.text.fill","stylers":[{"color":"#ffffff"}]},{"featureType":"road.highway","elementType":"labels.text.stroke","stylers":[{"color":"#ed5929"}]},{"featureType":"road.arterial","elementType":"geometry","stylers":[{"color":"#000000"},{"lightness":18}]},{"featureType":"road.arterial","elementType":"geometry.fill","stylers":[{"color":"#575757"}]},{"featureType":"road.arterial","elementType":"labels.text.fill","stylers":[{"color":"#ffffff"}]},{"featureType":"road.arterial","elementType":"labels.text.stroke","stylers":[{"color":"#2c2c2c"}]},{"featureType":"road.local","elementType":"geometry","stylers":[{"color":"#000000"},{"lightness":16}]},{"featureType":"road.local","elementType":"labels.text.fill","stylers":[{"color":"#999999"}]},{"featureType":"transit","elementType":"geometry","stylers":[{"color":"#000000"},{"lightness":19}]},{"featureType":"water","elementType":"geometry","stylers":[{"color":"#000000"},{"lightness":17}]}];
    this.positionInfoMarker = "fa-car";//fa-map-marker
    this.positionInfo = "La voiture est à l'arret"//La voiture est au garage
    this.apikey = apiService.getKey('maps');
    this.numberOfTrips = 5;
  }

  ngOnInit(){

    this.getTrips();
    this.getCardata();

    setInterval(() => {
      this.getTrips();
      this.getCardata();
    }, this.interval);

  }

  changeState(state){
    //garage
    //moveLille
    //moveOutside
    this.state = state;
  }

  centerMap(){
    //map.getBounds().contains(marker.getPosition())
  }

  getTrips(){
    this.carService.getCarData('trips')
    .map(function(res){
      for (let i = 0; i < res.length; i++) {
          res[i].beginDate = new Date(res[i].beginDate);
          res[i].stopDate = new Date(res[i].stopDate);
          var diffMins = Math.round((((res[i].stopDate - res[i].beginDate) % 86400000) % 3600000) / 60000); // minutes
          res[i].minutes = diffMins;
          res[i].distance = 0;
      }
      return res;
    }).subscribe((o) => {
      this.trips = [];
      if(this.numberOfTrips > o.length)
        this.numberOfTrips = o.length;
      for(let j = 1; j <= this.numberOfTrips; j++){
        this.trips.push(o[o.length-j]);
      }
      console.log(this.trips);
      //on recupere les distances des trajets
      //this.getTripsDistance();
    });
  }

  getTripsDistance(){
    var i = 0;
    Observable
      .from(this.trips)
      .flatMap((trip) => this.carService.getCarData('trip&trip_id='+trip.id))
      .subscribe((o) => {
        var signals = o;
        var min = 10000000;
        var max = 0;
        for (let k = 0; k < signals.length; k++) {
          if(signals[k].name === "Odometer"){
            if(signals[k].value < min){
              min = signals[k].value;
            }
            if(signals[k].value > max){
              max = signals[k].value;
            }
          }
        }
        this.trips[i].distance = max - min;
        i++;
      });

  }

  getCardata(){
    this.carService.getCarData('car')
    .subscribe((data) => {
      //location
      this.latCar = data.location.latitude;
      this.lngCar = data.location.longitude;
      //signals
      var signals = data.signals;
      this.lastUpdateDate = new Date(data.lastUpdateDate);
      for (let i = 0; i < signals.length; i++) {
          if(signals[i].name === "FuelLevel"){
            this.fuelLevel = signals[i].value;
            this.fuelRange = Math.round(this.fuelLevel/10);
          }else if(signals[i].name === "EngineSpeed"){
            this.engineSpeed = signals[i].value;
            if(this.engineSpeed > 0){
              this.positionInfo = "La voiture est en déplacement";
              this.positionClass = "move_me";
            }else{
              this.positionInfo = "La voiture est à l'arret";
              this.positionClass = "";
            }
          }
      }
    });
  }

}
