# importing ffmpeg streaming
import ffmpeg_streaming
from ffmpeg_streaming import Formats



# const
BASE = "tmp/dash/"

# getting file name
file_name = input("[NAME] > ")
file_id = input("[OUTPUT ID] > ")

# creating a video instance
video = ffmpeg_streaming.input(file_name)

print("Processing ...")

# creating a dash video
dash = video.dash(Formats.h264())
dash.auto_generate_representations()
dash.output(BASE + file_id + ".mpd")

print("Done.")
