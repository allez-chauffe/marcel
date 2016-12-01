import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { GithubService } from './github.service';

@Component({
  selector: 'github',
  encapsulation: ViewEncapsulation.None,
  templateUrl: './github.component.html',
  styleUrls: ['./github.component.scss'],
  providers : [ GithubService ]
})

export class GithubComponent implements OnInit {

  contributors : any[] = [];

  constructor(public githubService:GithubService) { }

  ngOnInit() {
    this.githubService.getTopContributors().subscribe((contributors)=>{
      contributors.sort(function(a,b){
        if (a.last_nom < b.last_nom)
          return -1;
        if (a.last_nom > b.last_nom)
          return 1;
        return 0;
      })
      this.contributors = contributors;
    });
  }
}
