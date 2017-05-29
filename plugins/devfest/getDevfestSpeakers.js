var items = document.getElementsByTagName('df-speaker')
var speakers = [];

for (var i in items) {
    var speaker = {};
    if (items[i].bio != undefined) {
        speaker.bio = items[i].bio;
        speaker.company = items[i].company;
        speaker.name = items[i].name;

        speakers.push(speaker);
    }
}


console.log(JSON.stringify(speakers));