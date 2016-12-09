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
    
  }

  ngOnInit() {
    if(!this.is_weekend()){
      this.lunchplaceService.get_teams_daily(this.organization).subscribe(orga => {
        console.log(orga);
        this.teams = orga.teams
      });
    }
  }

  is_weekend(){
    var day = new Date().getDay();
    return (day == 6) || (day == 0);
  }

}
