# importing ffmpeg streaming
import ffmpeg_streaming
from ffmpeg_streaming import Formats

import sys



# const
BASE = "dash/"

# getting file name
file_name = sys.argv[1]
file_id = sys.argv[2]

# creating a video instance
video = ffmpeg_streaming.input(file_name)

print("[OK] Processing ...")

# creating a dash video
dash = video.dash(Formats.h264())
dash.auto_generate_representations()
dash.output(BASE + file_id + ".mpd")

print("[OK] Done")
