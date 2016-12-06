import { Component, OnInit,ViewEncapsulation } from '@angular/core';

@Component({
  selector: 'anniversaire',
  encapsulation: ViewEncapsulation.None,
  templateUrl: './anniversaire.component.html',
  styleUrls: ['./anniversaire.component.scss']
})
export class AnniversaireComponent implements OnInit {

  users : any[] = [];

  constructor() { 
    this.users.push(
      {
        who:'Aurélien Loyer',
        date : '06/12/91',
        age : '43 ans'
      },
      {
        who:'Antoine Cordier',
        date : '06/12/91',
        age : '23 ans'
      }
    );
  }

  ngOnInit() {
  }

}
