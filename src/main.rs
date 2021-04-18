#![allow(unused_imports)]

extern crate ffmpeg_next as ffmpeg;

use std::io::Write;
use std::fs::File;
use ffmpeg::{
    codec, decoder, encoder, format, frame, log, media, picture, Dictionary, Packet, Rational,
};

use rustdct::DctPlanner;
use ffmpeg::format::Pixel;
use ffmpeg::media::Type;
use ffmpeg::software::scaling::{context::Context, flag::Flags};
use ffmpeg::util::frame::video::Video;

fn main() {
    ffmpeg::init().unwrap();

    let mut ictx = format::input(&"test-media/lenna.png").unwrap();
    println!("{:?}", ictx.format().name());
    let input = ictx.streams().best(Type::Video).ok_or(ffmpeg::Error::StreamNotFound).unwrap();
    let video_stream_index = input.index();
    let mut decoder = input.codec().decoder().video().unwrap();
    let mut scaler = Context::get(
        decoder.format(),
        decoder.width(),
        decoder.height(),
        Pixel::GRAY8,
        256,
        256,
        Flags::BILINEAR,
    ).unwrap();
    let mut frame_index = 0;

    let mut receive_and_process_decoded_frames =
        |decoder: &mut ffmpeg::decoder::Video| -> Result<(), ffmpeg::Error> {
            let mut decoded = Video::empty();
            while decoder.receive_frame(&mut decoded).is_ok() {
                let mut rgb_frame = Video::empty();
                scaler.run(&decoded, &mut rgb_frame)?;
                save_file(&rgb_frame, frame_index).unwrap();
                frame_index += 1;
            }
            Ok(())
        };

    for (stream, packet) in ictx.packets() {
        if stream.index() == video_stream_index {
            decoder.send_packet(&packet).unwrap();
            receive_and_process_decoded_frames(&mut decoder).unwrap();
        }
    }
    decoder.send_eof().unwrap();
    receive_and_process_decoded_frames(&mut decoder).unwrap();
}

fn save_file(frame: &Video, index: usize) -> std::result::Result<(), std::io::Error> {
    // println!("{}*{}={}", frame.width(), frame.height(), frame.data(0).len());
    let width = frame.width() as usize;
    let height = frame.height() as usize;
    let data_len = width * height;
    let mut planner = DctPlanner::new();
    let dct4_row = planner.plan_dct4(width as usize);
    let dct4_col = planner.plan_dct4(height as usize);

    // Convert data from u8 to f32 and make a few copies
    let mut data_row = Vec::<f32>::with_capacity(data_len);
    for pixel in frame.data(0) {
        data_row.push(*pixel as f32);
    }
    let mut data_col = data_row.clone();
    let mut data_final = data_row.clone();
    println!("{} {} {}", width, height, data_len);
    println!("{} {}", data_row.len(), data_col.len());

    // Compute DCT for rows
    for line in data_row.chunks_exact_mut(width) {
        dct4_row.process_dct4(line);
    }
    // Compute DCT for columns
    transpose::transpose(&data_row, &mut data_col, width, height);
    for line in data_col.chunks_exact_mut(height) {
        dct4_col.process_dct4(line);
    }
    // Multiply DCTs and normalize
    let data_len_f32 = data_len as f32;
    let norm_factor = (4.0/data_len_f32).sqrt();
    transpose::transpose(&data_col, &mut data_final, height, width);
    for i in 0..data_len {
        data_final[i] *= norm_factor;
        if i < width || i % width == 0 {
            data_final[i] /= 2.0_f32.sqrt();
        }
    }

    let mut nums = data_final.clone();
    nums.sort_by(|a, b| a.partial_cmp(b).unwrap());
    println!("{} {} {} {} {} {} {} {} {} {} {}", nums[0], nums[data_len/10], nums[2*data_len/10], nums[3*data_len/10], nums[4*data_len/10], nums[5*data_len/10], nums[6*data_len/10], nums[7*data_len/10], nums[8*data_len/10], nums[9*data_len/10], nums[data_len-1]);
    // println!("{} {} {} {} {}", nums[0], nums[1], nums[2], nums[3], nums[4]);
    let low_val = nums[data_len/10];
    let high_val = nums[9*data_len/10];

    // Convert back to u8
    let mut data_u8 : Vec<u8> = Vec::with_capacity(data_len);
    for pixel_f32 in &data_final {
        let mut val = *pixel_f32;
        // val *= (u8::MAX as f32) * (val - low_val)/(high_val - low_val);
        data_u8.push(val as u8);
    }
    // println!("{:?}", data_final);
    
    let mut file = File::create(format!("/tmp/eth21/frame{}.ppm", index))?;
    file.write_all(format!("P5\n{} {}\n255\n", frame.width(), frame.height()).as_bytes())?;
    // file.write_all(frame.data(0))?;
    file.write_all(&data_u8)?;

    Ok(())
}

fn save_file_old(frame: &Video, index: usize) -> std::result::Result<(), std::io::Error> {
    // println!("{}*{}={}", frame.width(), frame.height(), frame.data(0).len());
    let width = frame.width() as usize;
    let height = frame.height() as usize;
    let data_len = width * height;
    let mut planner = DctPlanner::new();
    let dct4_row = planner.plan_dct4(width as usize);
    let dct4_col = planner.plan_dct4(height as usize);

    // Convert data from u8 to f32 and make a few copies
    let mut data_row = Vec::<f32>::with_capacity(data_len);
    for pixel in frame.data(0) {
        // We use a scale from 0 to 1 to make things more compatible with OpenGL
        let mut pixel = *pixel as f32;
        pixel /= u8::MAX as f32;
        data_row.push(pixel);
    }
    let mut data_col = data_row.clone();
    let mut data_final = data_row.clone();

    // Compute DCT for rows
    for line in data_row.chunks_exact_mut(width) {
        dct4_row.process_dct4(line);
    }
    // Compute DCT for columns
    transpose::transpose(&data_row, &mut data_col, width, height);
    for line in data_col.chunks_exact_mut(height) {
        dct4_col.process_dct4(line);
    }
    // Multiply DCTs and normalize
    let data_len_f32 = data_len as f32;
    let norm_factor = (4.0/data_len_f32).sqrt();
    transpose::transpose(&data_col, &mut data_final, height, width);
    for i in 0..data_len {
        data_final[i] *= norm_factor*data_row[i];
    }

    let mut nums = data_final.clone();
    nums.sort_by(|a, b| a.partial_cmp(b).unwrap());
    println!("{} {} {} {} {} {} {} {} {} {} {}", nums[0], nums[data_len/10], nums[2*data_len/10], nums[3*data_len/10], nums[4*data_len/10], nums[5*data_len/10], nums[6*data_len/10], nums[7*data_len/10], nums[8*data_len/10], nums[9*data_len/10], nums[data_len-1]);
    // println!("{} {} {} {} {}", nums[0], nums[1], nums[2], nums[3], nums[4]);
    let low_val = nums[data_len/10];
    let high_val = nums[9*data_len/10];

    // Convert back to u8
    let mut data_u8 : Vec<u8> = Vec::with_capacity(data_len);
    for pixel_f32 in &data_final {
        let mut val = *pixel_f32;
        // val *= (u8::MAX as f32) * (val - low_val)/(high_val - low_val);
        val *= u8::MAX as f32;
        data_u8.push(val as u8);
    }
    // println!("{:?}", data_final);
    
    let mut file = File::create(format!("/tmp/eth21/frame{}.ppm", index))?;
    file.write_all(format!("P5\n{} {}\n255\n", frame.width(), frame.height()).as_bytes())?;
    // file.write_all(frame.data(0))?;
    file.write_all(&data_u8)?;

    // TODO: convert back to see if the image is being preserved
    // let mut revdata_f32 = Vec::<f32>::with_capacity(data_len);
    // for pixel in data_u8 {
    //     let mut pixel = *pixel as f32;
    //     pixel /= u8::MAX as f32;
    //     revdata_f32.push(pixel);
    // }

    

    Ok(())
}