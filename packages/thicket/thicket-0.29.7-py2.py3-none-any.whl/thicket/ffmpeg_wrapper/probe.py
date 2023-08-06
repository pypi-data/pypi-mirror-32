#!/usr/bin/env python
import json
import subprocess

from thicket.ffmpeg_wrapper.utils import is_video

try:
    from subprocess import DEVNULL
except ImportError:
    import os
    DEVNULL = open(os.devnull, 'wb')


def video_length(input_video):

    if is_video(input_video):
        info = video_info(input_video)
        format_length = info.get('format').get('duration')
        video_stream_length = 0
        audio_stream_length = 0

        for stream in info.get('streams'):
            if stream.get('codec_type') == 'video':
                video_stream_length = stream.get('duration')
            if stream.get('codec_type') == 'audio':
                audio_stream_length = stream.get('duration')

        return dict(
            video_stream_length=video_stream_length,
            audio_stream_length=audio_stream_length,
            format_length=format_length
        )

    default = float(0)
    return dict(
        video_stream_length=default,
        audio_stream_length=default,
        format_length=default
    )


def video_info(input_video):

    if is_video(input_video):

        command_list = ['/usr/bin/ffprobe',
                        '-i',
                        input_video,
                        '-show_format',
                        '-show_streams',
                        '-print_format',
                        'json']

        stdout, stderr = subprocess.Popen(command_list, stdout=subprocess.PIPE, stderr=subprocess.PIPE).communicate()
        strout = stdout.decode('utf-8')
        data = json.loads(strout)
        return data

    default = {}
    return default
