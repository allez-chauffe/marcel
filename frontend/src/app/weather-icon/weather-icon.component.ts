import {Component, OnInit,Input} from "@angular/core";
import {ViewEncapsulation} from '@angular/core';

import {WeatherService} from '../weather/weather.service';
import {FirstletterupperPipe} from '../pipes/firstletterupper.pipe';
import { WeatherIconService } from './weather-icon.service';

@Component({
  selector: 'weather-icon',
  encapsulation: ViewEncapsulation.None,
  templateUrl: './weather-icon.component.html',
  styleUrls: ['./weather-icon.component.scss'],
  providers : [WeatherIconService]
})

export class WeatherIconComponent implements OnInit{

  @Input() numberlist : number;
  @Input() displaydt : boolean;

  public current_weather_type : String = "";
  public weather_types : Array<Object> = [];
  public weather_object : any;
  public weather_date : Date;
  public weather_temp : number = 0 ;
  public wheather_icon : string;

  private timer: number = 180000;
  
  constructor(public weatherService: WeatherService, public weatherIconService : WeatherIconService){
    this.weather_types.push("sun-shower");
    this.weather_types.push("thunder-storm");
    this.weather_types.push("cloudy");
    this.weather_types.push("sunny");
    this.weather_types.push("rainy");
    this.weather_types.push("flurries");
  }

  ngOnInit(){
    console.log('Init weather icon');
    this.getWeather();
    setInterval(() => {
      this.getWeather();
    }, this.timer);
  }

  getWeather(){
    this.weatherService.getCurrent()
    .subscribe(res => {
      this.weather_temp = Math.round(res.list[this.numberlist].main.temp);
      this.setWeatherType(res.list[this.numberlist].weather[0].main);
      this.weather_date = new Date(res.list[this.numberlist].dt * 1000);
      this.wheather_icon = this.getDayNight(res.list[this.numberlist].dt * 1000) + this.weatherIconService.getOneIcon(res.list[this.numberlist].weather[0].id).icon;
    });
  }

  randomType(){
    var randomInt = Math.floor(Math.random() * this.weather_types.length);
    return this.weather_types[randomInt];
  }

  getDayNight(date){
    let hr = (new Date(date)).getHours();
    console.log(hr);
    if(hr > 6 && hr < 18)
      return 'day-';
    else
      return 'day-';//night
  }

  setWeatherType(type) {
    if(!type){
      type = this.randomType();
    }
    this.current_weather_type = type;
  }

}
