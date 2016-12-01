import { Component, OnInit,ViewEncapsulation } from '@angular/core';

@Component({
  selector: 'lunchplace',
  encapsulation: ViewEncapsulation.None,
  templateUrl: './lunchplace.component.html',
  styleUrls: ['./lunchplace.component.scss']
})
export class LunchplaceComponent implements OnInit {

  teams : any[] = [];

  constructor() {
    this.teams.push(
      {
        name:'Les biloutes',
        restaurant : 'Le Tokyo',
        icon : 'http://image.flaticon.com/icons/svg/68/68273.svg'
      },
      {
        name:'Les stagiaires',
        restaurant : 'McDonald',
        icon : 'cutlery'
      }
    );
  }

  ngOnInit() {
  }

}
