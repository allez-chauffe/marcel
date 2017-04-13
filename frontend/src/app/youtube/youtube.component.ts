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
  id: String = '';
  player: any;

  constructor(private youtubeService: YoutubeService) {
    this.subscription = this.youtubeService.getSearch().subscribe(message => this.search(message.query));
  }

  ngOnInit() {
  }

  savePlayer(player) {
    this.player = player;
    this.player.playVideo();
  }

  search(query) {
    this.youtubeService.search(query).subscribe(videos => {
      this.videos = videos;
      this.id = videos[0].id.videoId;
      if(this.player !== undefined){
        this.player.loadVideoById(this.id);
        this.player.playVideo();
      }
    });
  }

}
