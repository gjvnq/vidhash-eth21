import cv2
import imagehash
import numpy as np
from PIL import Image
from math import sqrt
from skimage.color import rgb2gray
from bitstring import BitArray, BitStream
from skimage.transform import resize as imresize

# implement 2D DCT
def dct2(img):
	from scipy.fftpack import dct
	return dct(dct(img.T, norm='ortho').T, norm='ortho')

# implement 2D IDCT
def idct2(img):
	from scipy.fftpack import idct
	return idct(idct(img.T, norm='ortho').T, norm='ortho')   

def diag_gen(nsides):
	x, y = 0, 0
	upright = True
	visits = 1
	while visits <= nsides**2:
		yield (x,y)
		visits += 1
		if upright:
			x, y = x - 1, y + 1
		else:
			x, y = x + 1, y - 1
		#yield (0, x, y)
		if 0 <= x < nsides:
			if 0 <= y < nsides:
				continue
			elif 0 < y:
				x += 2
				y -= 1
			else:
				y = 0
		elif 0 < x:
			if 0 <= y < nsides:
				x -= 1
				y += 2
			elif 0 < y:
				raise Exception("this shouldn't happen ({}, {})".format(x, y))
			else:
				x -= 1
				y += 2
		else:
			if 0 <= y < nsides:
				x = 0
			elif 0 < y:
				x += 2
				y = nsides-1
			else:
				raise Exception("this shouldn't happen ({}, {})".format(x, y))
		upright = not upright

# must be already grayscale
def img_to_hash(img_full, resizeTo=64, cropTo=16, plot=False):
	#img_full = rgb2gray(imread(src))
	img_base = imresize(img_full, (resizeTo,resizeTo))
	img_base -= 0.5
	
	img_freq = dct2(img_base)
	amortize = np.vectorize(lambda x: np.sign(x)*sqrt(abs(x)))
	img_freq_crop = img_freq[0:cropTo,0:cropTo]
	img_freq_amort = amortize(img_freq_crop)
	
	freq_avg = img_freq_amort.mean()
	f = lambda x: x > freq_avg
	img_freq_final = f(img_freq_amort)
	
	hash_val = "0b"
	for (x, y) in diag_gen(cropTo):
		hash_val += str(int(img_freq_final[x][y]))
	hash_val = BitArray(hash_val)
	return hash_val

def hamming_dist(a,b):
	"""Calculate the Hamming distance between two bit strings"""
	assert len(a.bin) == len(b.bin)
	count = 0
	return (a^b).count(1) #XOR




def step1(videopath):
	vidcap = cv2.VideoCapture(videopath)
	success,image = vidcap.read()
	if not success:
		raise Exception("failed to open video")
	fps = vidcap.get(cv2.CAP_PROP_FPS)

	print("Video FPS: {}".format(fps))
	# print("\033[48;2;255;140;60m ORANGE BACKGROUND \033[48;2;0;0;0m")

	last_hash = None
	frame_len = int(vidcap.get(cv2.CAP_PROP_FRAME_COUNT))
	frames = []
	frames_hash_diff = [0]
	while success:
		success, image = vidcap.read()
		if success:
			image = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)
			frames.append(image)
			hash_val = img_to_hash(image)
			if last_hash is not None:
				frames_hash_diff.append(hamming_dist(last_hash, hash_val))
			last_hash = hash_val
		if (len(frames)-1) % 100*fps == 0:
			percentage = round(100*len(frames)/frame_len)
			print("\033[0K\033[0G"+f"Processed {len(frames)} frames ({percentage}%)", end="", flush=True)

	print(f"\033[0K\033[0G"+f"Processed {len(frames)} frames (100%)", flush=True)
	print(f"Video frame length: {len(frames)}")
	print(f"Video time length: {round(len(frames)/fps, 1)} s")

	return frames, frames_hash_diff, fps

def segment_video(frames, frames_hash_diff_2, fps):
	segments = []
	start = None
	end = None
	min_scene_length = 3*fps
	for frame, diff in enumerate(frames_hash_diff_2):
		if diff <= 20:
			if start is None:
				start = frame
			else:
				end = frame
		else:
			if start is not None:
				if end-start >= min_scene_length:
					segments.append((start, end))
				start = None
	return segments

def hash_segment(frames, segment):
	start, end = segment
	avg_frame = Image.fromarray(np.average(frames[start:end], axis=0))
	return imagehash.phash(avg_frame)

def vp_hash_file(videopath):
	frames, frames_hash_diff, fps = step1(videopath)
	frames_hash_diff_2 = [0] + [frames_hash_diff[i-1]-frames_hash_diff[i] for i in range(1, len(frames_hash_diff))]
	segments = segment_video(frames, frames_hash_diff_2, fps)
	key_segments_hashes = []
	for segment in segments:
		key_segments_hashes.append(hash_segment(frames, segment))
	return key_segments_hashes