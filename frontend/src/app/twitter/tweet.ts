import { TwitterUser } from './twitter-user';

export class Tweet {

	constructor(public created_at: Date, public text: string,
				public truncated: boolean, public user: TwitterUser){ }
}
