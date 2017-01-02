import { SoundtouchService } from './soundtouch.service';
import { SoundtouchComponent } from './soundtouch.component';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

@NgModule({
    declarations: [ SoundtouchComponent ],
    exports: [ SoundtouchComponent ],
    providers: [ SoundtouchService],
    imports: [ CommonModule ]
})
export class SoundtouchModule {

}