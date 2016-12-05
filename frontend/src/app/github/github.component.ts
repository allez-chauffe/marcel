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

  constructor(public githubService: GithubService) { }

  ngOnInit() {
    setInterval(() => {
      this.fetchContributors();
    }, 1000 * 60 * 60);
    this.fetchContributors();
  }

  fetchContributors(){
    this.githubService.getTopContributors().subscribe(contributors => this.contributors = contributors);
  }
}
