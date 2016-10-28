import {Component, OnInit,ElementRef} from "@angular/core";
import {ViewEncapsulation} from '@angular/core';
import {NgStyle} from '@angular/common';
import {AddZeroPipe} from '../pipes/addzero.pipe';

@Component({
  selector: 'datetime',
  encapsulation: ViewEncapsulation.None,
  templateUrl: './datetime.component.html',
  styleUrls: ['./datetime.component.scss']
})

export class DateTimeComponent implements OnInit{

  public numDay: number;
  public numMonth: number;
  public numYear: number;
  public strMonth: string;
  public seconds:number;
  public minutes:number;
  public hour:number;

  constructor() {
  }

  ngOnInit(){
    console.log('Init DateTime');
    this.loopDate();
    setInterval(() => {
      this.loopDate();
    }, 10000);
  }

  loopDate(){
    var date = new Date();
    var tab_mois = new Array("Janvier", "Février", "Mars", "Avril", "Mai", "Juin", "Juillet", "Août", "Septembre", "Octobre", "Novembre", "Décembre");

    this.seconds = date.getSeconds();
    this.minutes = date.getMinutes();
    this.hour = date.getHours();

    this.numDay = date.getDate();
    this.numMonth = date.getMonth()+1;
    this.numYear = date.getFullYear();
    this.strMonth = tab_mois[date.getMonth()];
  }
}
