const prism = require("prism-media");

const encoder = new prism.opus.Encoder({ 
    rate: 48000, 
    channels: 2, 
    frameSize: 960
});

// Convert raw PCM to Opus
const transform = new prism.opus.OggDemuxer();
process.stdin.pipe(transform).pipe(encoder).pipe(process.stdout);
