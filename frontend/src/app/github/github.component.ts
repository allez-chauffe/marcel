import { Observable } from 'rxjs/Rx';
import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { GithubService } from './github.service';

@Component({
  selector: 'marcel-github',
  encapsulation: ViewEncapsulation.None,
  templateUrl: './github.component.html',
  styleUrls: ['./github.component.scss']
})

export class GithubComponent implements OnInit {

  contributors: any[];

  private timer: number = 1000 * 60 * 60;

  constructor(public githubService: GithubService) { }

  ngOnInit() {
    setInterval(() => {
      this.fetchContributors();
    }, this.timer);
    this.fetchContributors();
  }

  fetchContributors(){
    this.githubService.getTopContributors().subscribe(contributors => this.contributors = contributors);
  }
}
