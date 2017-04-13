import { Component, OnInit, Input } from '@angular/core';
import { Subscription } from 'rxjs/Subscription';
import { YoutubeService } from './youtube.service';

@Component({
  selector: 'youtube',
  templateUrl: './youtube.component.html',
  styleUrls: ['./youtube.component.scss']
})
export class YoutubeComponent implements OnInit {
  videos: Array<any>;
  subscription: Subscription;

  constructor(private youtubeService: YoutubeService) {
    this.subscription = this.youtubeService.getSearch().subscribe(message => this.search(message.query));
  }

  ngOnInit() {
  }

  search(query) {
    this.youtubeService.search(query).subscribe(videos => {
      this.videos = videos;
      console.log(videos);
      console.log(videos[0].id.videoId);
    });
  }

}
