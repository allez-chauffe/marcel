import { YunService } from './yun.service';
import { HomeComponent } from './home.component';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

@NgModule({
    declarations: [ HomeComponent ],
    exports: [ HomeComponent ],
    providers: [ YunService ],
    imports: [ CommonModule ]
})
export class HomeModule {

}