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
  searchSubscription: Subscription;
  fullScreenSubscription: Subscription;
  id: String = '';
  player: any;

  constructor(private youtubeService: YoutubeService) {
    this.searchSubscription = this.youtubeService.getSearch().subscribe(message => this.search(message.query));
    this.fullScreenSubscription = this.youtubeService.getQuery().subscribe(message => this.handleQuery(message.type));
  }

  ngOnInit() {
  }

  launchIntoFullscreen(element: any) {

  }

  handleQuery(type) {
    if(type === "fullscreen"){
      var element = document.getElementById("youtube-player").getElementsByTagName("iframe")[0] as any;
      if(element.requestFullscreen) {
        element.requestFullscreen();
      } else if(element.mozRequestFullScreen) {
        element.mozRequestFullScreen();
      } else if(element.webkitRequestFullscreen) {
        element.webkitRequestFullscreen();
      } else if(element.msRequestFullscreen) {
        element.msRequestFullscreen();
      }
    }

    if(type === "pause"){
      this.player.pauseVideo();
    }

    if(type === "play"){
      this.player.playVideo();
    }
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
