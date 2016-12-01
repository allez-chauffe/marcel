import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { GithubService } from './github.service';

@Component({
  selector: 'github',
  encapsulation: ViewEncapsulation.None,
  templateUrl: './github.component.html',
  styleUrls: ['./github.component.scss']
})

export class GithubComponent implements OnInit {

  contributors: any[];

  constructor(public githubService: GithubService) { }

  ngOnInit() {
    this.githubService.getTopContributors().subscribe(contributors => this.contributors = contributors)
  }
}
