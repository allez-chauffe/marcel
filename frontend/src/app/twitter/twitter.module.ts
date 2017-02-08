import { TwitterService } from './twitter.service';
import { TwitterComponent } from './twitter.component';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

@NgModule({
    declarations: [ TwitterComponent ],
    exports: [ TwitterComponent ],
    providers: [ TwitterService ],
    imports: [ CommonModule ]
})
export class TwitterModule {

}
