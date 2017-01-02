import { Component, OnInit,ElementRef } from "@angular/core";
import { ViewEncapsulation } from '@angular/core';
import { NgStyle } from '@angular/common';
import { ApiService } from '../api/api.service';
import { YunService } from './yun.service';

@Component({
  selector: 'home',
  encapsulation: ViewEncapsulation.None,
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})

export class HomeComponent implements OnInit{

  public _sRotate: string = "1";
  public rooms : any[] = [];

  private interval: number = 1000 * 60 * 60;

  constructor(public yunService: YunService) {
  }

  ngOnInit(){
    setInterval(() => {
      this.getRoomsStats();
    }, this.interval);
    this.getRoomsStats();
  }

  getRoomsStats(){
    this.yunService.getRoomStat('sparrow')
      .subscribe(stat => {
        this.rooms.push({
          name : 'Salle Sparrow',
          stat : stat
        });
      });

    this.yunService.getRoomStat('smaug')
      .subscribe(stat => {
        this.rooms.push({
          name : 'Salle Smaug',
          stat : stat
        });
      });
  }
}
