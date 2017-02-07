import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { TwitterService } from './twitter.service';
import { Tweet } from './tweet';

@Component({
  selector: 'marcel-twitter',
  encapsulation: ViewEncapsulation.None,
  templateUrl: './twitter.component.html',
  styleUrls: ['./twitter.component.scss']
})
export class TwitterComponent implements OnInit {

  private tweets: Tweet[];
  private timer: number = 1000 * 60 * 60;

  constructor(private twitterService: TwitterService) { }

  ngOnInit() {
    this.getTimeline();
    setInterval(this.getTimeline, this.timer);
  }

  getTimeline() {
    this.twitterService.getTimeline(1).subscribe(tweets => this.tweets = tweets);
   }

}
