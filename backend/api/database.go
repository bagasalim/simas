package api

import (
	"fmt"
	"os"

	"github.com/bagasalim/simas/model"
	// "github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDb() (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	dbUrl := os.Getenv("DATABASE_URL")

	if os.Getenv("ENVIRONMENT") == "PROD" {
		db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	} else {
		host := os.Getenv("host")
		port := os.Getenv("port_db")
		user := os.Getenv("user")
		password := os.Getenv("password")
		dbname := os.Getenv("dbname")
		config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

		db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	if os.Getenv("AUTO_MIGRATE") == "Y" {
		if err := db.AutoMigrate(model.User{}, model.Link{}, model.Riwayat{}, model.InfoPromo{}, model.Asuransi{}); err != nil {
			return nil, fmt.Errorf("failed to migrate database: %w", err)
		}

		users := []model.User{
			{
				Username: "admin",
				Password: "$2a$10$BQHCjmHmEsFGJXCGWm7et.2lvVPecg0ibhFd/tgOCCCncTu5ieiA.",
				Name:     "Administrator",
				Role:     1,
			},
			{
				Username: "CS01",
				Password: "$2a$10$BQHCjmHmEsFGJXCGWm7et.2lvVPecg0ibhFd/tgOCCCncTu5ieiA.",
				Name:     "Customer Service",
				Role:     2,
			},
		}

		links := []model.Link{
			{
				LinkType:  "WA",
				LinkValue: "https://api.whatsapp.com/send?phone=6288221500153",
				UpdatedBy: "System",
			},
			{
				LinkType:  "Zoom",
				LinkValue: "https://zoom.us/w/99582712162?tk=-ZsgZOP5esSZvy2g1sfWt8R3ugl9woAjQGuFFgUaU3k.DQMAAAAXL5eZYhZvdW5zcWJ4elJvaUt3cHFza1FBaVZRAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA&pwd=SzRUOFNIVldlRkR6SlFpc004OUs1Zz09",
				UpdatedBy: "System",
			},
		}

		riwayats := []model.Riwayat{
			{
				Nama:       "John",
				Email:      "john@gmail.com",
				Kategori:   "Kartu Kredit",
				Keterangan: "Complain CC",
			},
			{
				Nama:       "Doe",
				Email:      "doe@gmail.com",
				Kategori:   "Digital Loan",
				Keterangan: "Cara Daftar Loan",
			},
		}

		infopromos := []model.InfoPromo{
			{
				Judul:     "Gebyar Sinarmas",
				Kategori:  "Promo Simobiplus",
				Startdate: "2022-10-10",
				Enddate:   "2022-10-30",
				Kodepromo: "202223",
				Foto:      "https://webpraktis.com/medias/blog/z9gkkk8dhx.png",
				Deskripsi: "Gebyar sinarmas hadir untuk memeriahkan hari kemerdekaan indonesia, ayo join dan gebyarkan indonesia bersama sinarmas dan nikmati keunggulan diskon pembayaran melalui simobiplus",
				Syarat:    "1. Satu rekening hanya bisa melakukan pembayaran satu kali; 2. Satu nomor hp hanya bisa melakukan pembayaran satu kali; 3. Nasabah dapat membuka rekening melalui simobiplus",
			},
			{
				Judul:     "Belanja di Alfamart & UniPin",
				Kategori:  "Promo Kartu Kredit",
				Startdate: "2022-04-11",
				Enddate:   "2022-06-05",
				Kodepromo: "202123",
				Foto:      "https://glints.com/id/lowongan/wp-content/uploads/2019/12/clothes-on-sale-2292953.jpg",
				Deskripsi: "Cukup belanja Rp100 ribu di Alfamart atau UniPin melalui fitur Lifetyle SimobiPlus, kamu BISA dapat hadiah E-Voucher Perfect Beauty senilai Rp100 ribu! Nilai transaksi berlaku akumulasi selama periode program berlangsung.",
				Syarat:    "1. Nasabah harus melakukan transaksi senilai Rp100 ribu di Alfamart atau UniPin melalui Fitur Lifestyle SimobiPlus; 2. Perhitungan belanja/transaksi Rp 100 ribu berlaku akumulasi selama program berlangsung; 3. Hadiah E-Voucher Perfect Beauty akan diterima Nasabah melalui aplikasi SimobiPlus pada 21 Oktober 2022",
			},
			{
				Judul:     "Penuhi Gaya Hidupmu Lebih Ringan dengan Cicilan 0% Tenor 3 Bulan via SimobiPlus",
				Kategori:  "Promo Bank",
				Startdate: "2022-10-11",
				Enddate:   "2022-12-25",
				Kodepromo: "102121",
				Foto:      "https://glints.com/id/lowongan/wp-content/uploads/2019/12/clothes-on-sale-2292953.jpg",
				Deskripsi: "Belanja, traveling, kulineran, atau penuhi kebutuhan lainnya tak perlu ragu pakai Kartu Kredit Platinum Bank Sinarmas. Nikmati cicilannya sekarang via SimobiPlus!",
				Syarat:    "1. Transaksi yang telah mendapatkan potongan langsung/diskon tidak dapat dijadikan cicilan pada aplikasi merchant E-commerce; 2. Transaksi dapat dibatalkan apabila barang rusak atau tidak tersedia/out of stock, dengan catatan transaksi dibatalkan sebelum tagihan jatuh tempo. Refund yang masuk atas pembatalan transaksi ini akan secara otomatis mengurangi tagihan bulan berjalan; 3. Program Reguler tidak berlaku jika telah memasuki Program Payday.",
			},
		}
		asuransis := []model.Asuransi{
			{
				Judul:             "Simas Insurtech Travel",
				Premi:             200000,
				UangPertanggungan: 100000000,
				Deskripsi:         "Memberikan penggantian yang maksimal sesuai dengan program yang dipilih jika Tertanggung meninggal dunia/cacat tetap akibat kecelakaan",
				Syarat:            "Minimal 17 tahun dan maksimal 62 tahun, WNI",
				Foto:              "https://www.banksinarmas.com/id/public/upload/thumb/60b8a9ec8489a_thumb-simas-insurtech-travel.jpg",
			},
			{
				Judul:             "Simas Pet Insurance",
				Premi:             200000,
				UangPertanggungan: 100067000,
				Deskripsi:         "Jalani aktivitas seru dengan kenyamanan dan perlindungan dari Simas Pet Insurance. Sepenuh hati #UntukYangTersayang",
				Syarat:            "Minimal 17 tahun dan maksimal 62 tahun, WNI",
				Foto:              "https://img11.jd.id/Indonesia/s380x380_/amZzL3Q4Mi8xMjQvMTA3MDkyODg5OTQvNjc4Mjk2L2VmNWExZTg2LzYxMjcxYTQ3Tjc3NGM1ZTE4.jpg.dpg.webp",
			},
		}

		resUsers := db.Create(&users)
		if resUsers == nil {
			return nil, fmt.Errorf("failed to seeding users database: %w", resUsers.Error)
		}

		resLinks := db.Create(&links)
		if resLinks == nil {
			return nil, fmt.Errorf("failed to seeding links database: %w", resLinks.Error)
		}

		resRiwayats := db.Create(&riwayats)
		if resRiwayats == nil {
			return nil, fmt.Errorf("failed to seeding riwayats database: %w", resRiwayats.Error)
		}

		resInfoPromos := db.Create(&infopromos)
		if resInfoPromos == nil {
			return nil, fmt.Errorf("failed to seeding info promos database: %w", resInfoPromos.Error)
		}
		resAsuransis := db.Create(&asuransis)
		if resAsuransis == nil {
			return nil, fmt.Errorf("failed to seeding asuransi database: %w", resAsuransis.Error)
		}

	}

	return db, err
}
