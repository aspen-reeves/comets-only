import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Profile } from './profile.model';
import { Profile2 } from './profile2.model';

@Injectable({ providedIn: 'root' })
export class APIService {
  profile: Profile2;

  constructor(private httpClient: HttpClient) {}

  getProfile() {
    this.httpClient
      .get<Profile2>('http://144.126.154.126:10000/getbitches')
      .subscribe((data: Profile2) => {
        console.log(data);
        this.profile = {...data};
      });
  }
}
