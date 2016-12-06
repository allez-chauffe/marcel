import { Component, OnInit,ViewEncapsulation,Input } from '@angular/core';
import { LunchplaceService } from './lunchplace.service';

@Component({
  selector: 'lunchplace',
  encapsulation: ViewEncapsulation.None,
  templateUrl: './lunchplace.component.html',
  styleUrls: ['./lunchplace.component.scss']
})
export class LunchplaceComponent implements OnInit {

  teams : any[] = [];
  @Input() organization : any;


  constructor(private lunchplaceService:LunchplaceService) {
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
    this.lunchplaceService.get_teams_by_orga(this.organization).subscribe(test => {
      console.log(test);
    });
  }

}
