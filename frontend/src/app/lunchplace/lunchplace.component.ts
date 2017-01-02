import { Component, OnInit, ViewEncapsulation, Input } from '@angular/core';
import { LunchplaceService } from './lunchplace.service';

@Component({
  selector: 'lunchplace',
  encapsulation: ViewEncapsulation.None,
  templateUrl: './lunchplace.component.html',
  styleUrls: ['./lunchplace.component.scss']
})
export class LunchplaceComponent implements OnInit {

  @Input() organization: any;
  
  teams: any[] = [];
  
  private timer: number = 1000 * 60 * 15;

  constructor(private lunchplaceService: LunchplaceService) { }

  ngOnInit() {
    setInterval(() => {
      this.lunchplaceService
        .get_teams_daily(this.organization)
        .subscribe(orga => {
          this.teams = orga.teams
        });
    }, this.timer);
  }

}
