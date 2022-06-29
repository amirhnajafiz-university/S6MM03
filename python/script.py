# importing ffmpeg streaming
import ffmpeg_streaming
from ffmpeg_streaming import Formats



# const
BASE = "tmp/dash/"

# getting file name
file_name = input()
file_id = input()

# creating a video instance
video = ffmpeg_streaming.input(file_name)

# creating a dash video
dash = video.dash(Formats.h264())
dash.auto_generate_representations()
dash.output(BASE + file_id)
