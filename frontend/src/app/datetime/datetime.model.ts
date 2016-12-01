export class DateTime {
    public day: string;
    public year: string;
    public month: string;
    public seconds: string;
    public minutes: string;
    public hour: string;

    constructor(date: Date) {
        this.seconds = date.toLocaleString('fr', { 'second': '2-digit' })
        this.minutes = date.toLocaleString('fr', { 'minute': '2-digit' })
        this.hour = date.toLocaleString('fr', { 'hour': '2-digit' })

        this.day = date.toLocaleString('fr', { 'day': '2-digit' });
        this.year = date.toLocaleString('fr', { 'year': 'numeric' });
        this.month = date.toLocaleString('fr', { 'month': 'long' });
    }
}