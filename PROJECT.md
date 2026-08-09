# PROJECT.md — Project Brief

## What this app does
Take music from YouTube, or an imported audio file, and convert it into code usable by QBASIC or Makecode Adafruit.

## Roadmap for v1
- [x] 0. Skeleton app with empty window
- [x] 1. Input can be a YouTube URL. Example, to use for testing: https://www.youtube.com/watch?v=siCmqvfw_1g (this is No Copyright Music free to use or modify).
- [x] 2. Alternatively, input can be a wav, m4u or mp3 audio file. Example, to use for testing: C:\source\music-basicifier\example-input\Useless-Station.wav (copyright owned).
- [x] 3. Once input is given, user confirms with a button click or "Enter" keypress.
- [x] 4. If input is a YouTube URL, it obtains the audio from it. Use yt-dlp library to extract the audio.
- [ ] 5. Open the audio file and extract the main melody from it.
  - [ ] 5.1. Accept a local audio file path and validate that the file exists and is readable.
  - [ ] 5.2. Decode the audio file and isolate a single dominant melody stream into a temporary representation.
  - [ ] 5.3. Expose the extracted melody data to the downstream conversion layer with clear error handling for unsupported or unreadable files.
- [ ] 6. Convert the melody into 2 text outputs - 
- [ ] 7. The first text output will be a QBASIC program that will perform a close imitation using the PLAY and SOUND commands
- [ ] 8. The second text output will be JavaScript compatible with Makecode Adafruit. An example output for Adafruit can be found at C:\source\music-basicifier\example-output\chariots-of-fire.js
- [ ] 9. Text outputs can be easily copied to the Windows clipboard.
- [ ] 10. Clear error messages logged if the audio can't be extracted or some other error occurs. Simple UI error message displayed to user.

## Explicitly out of scope for v1
- Only extract a single melody line from the audio file. v2 will enable multiple lines to be extracted (melody, harmony, tenor, bass, percussion, etc.)
- Only QBASIC version of BASIC for now. And no special tricks with direct memory manipulation, such as POKE, OUT, etc.
- No need to display progress.

## Architecture decisions already made
- GUI: Fyne v2 (pinned in go.mod) — do not add Wails/Walk/etc.
- Persistence: None required besides error log. Textual output only.
- Concurrency model: Single main goroutine + fyne's own event loop;
  Processing of audio files should be in a background thread, so UI remains responsive.
- Target: Windows 10/11 only, amd64. No cross-platform requirement even
  though Fyne supports it.

## Naming/domain vocabulary
[Any terms specific to your problem domain the agent should use
consistently — e.g. "Job" not "Task", "Profile" not "User".]

## Non-functional requirements
- Must run without admin rights: yes