// Binary ztex_info retrieves information from ZTEX USB devices.
package main

import (
	"fmt"
	"log"

	"github.com/aljumi/ztex"
	"github.com/google/gousb"

	getopt "github.com/pborman/getopt/v2"
)

var (
	allFlag = getopt.BoolLong("all", 'a', "output all information")

	usbFlag   = getopt.BoolLong("usb", 'u', "output USB device information")
	ztexFlag  = getopt.BoolLong("ztex", 'z', "output ZTEX device information")
	fpgaFlag  = getopt.BoolLong("fpga", 'f', "output FPGA status information")
	flashFlag = getopt.BoolLong("flash", 's', "output flash status information")

	helpFlag = getopt.BoolLong("help", 'h', "display this help and exit")
)

func printUSB(d *gousb.Device) error {
	mfr, err := d.Manufacturer()
	if err != nil {
		return fmt.Errorf("(*gousb.Device).Manufacturer: %v", err)
	}
	prd, err := d.Product()
	if err != nil {
		return fmt.Errorf("(*gousb.Device).Product: %v", err)
	}
	snr, err := d.SerialNumber()
	if err != nil {
		return fmt.Errorf("(*gousb.Device).SerialNumber: %v", err)
	}

	fmt.Printf("USB Device:\n")
	fmt.Printf("  Descriptor:\n")
	fmt.Printf("    Bus: %v\n", d.Desc.Bus)
	fmt.Printf("    Address: %v\n", d.Desc.Address)
	fmt.Printf("    Speed: %v\n", d.Desc.Speed)
	fmt.Printf("    Port: %v\n", d.Desc.Port)
	fmt.Printf("    Spec: %v\n", d.Desc.Spec)
	fmt.Printf("    Device: %v\n", d.Desc.Device)
	fmt.Printf("    Vendor ID: %v\n", d.Desc.Vendor)
	fmt.Printf("    Product ID: %v\n", d.Desc.Product)
	fmt.Printf("    Class: %v\n", d.Desc.Class)
	fmt.Printf("    Subclass: %v\n", d.Desc.SubClass)
	fmt.Printf("    Protocol: %v\n", d.Desc.Protocol)
	fmt.Printf("    Maximum Control Packet Size: %v\n", d.Desc.MaxControlPacketSize)
	fmt.Printf("  Manufacturer: %v\n", mfr)
	fmt.Printf("  Product: %v\n", prd)
	fmt.Printf("  Serial Number: %v\n", snr)

	return nil
}

func printZTEX(d *ztex.Device) error {
	fmt.Printf("ZTEX Device:\n")
	fmt.Printf("  Descriptor:\n")
	fmt.Printf("    Size: %v\n", d.DescriptorSize)
	fmt.Printf("    Version: %v\n", d.DescriptorVersion)
	fmt.Printf("    Magic: %v\n", d.DescriptorMagic)
	fmt.Printf("    Product ID: %v\n", d.DescriptorProduct)
	fmt.Printf("    Firmware Version: %v\n", d.DescriptorFirmware)
	fmt.Printf("    Interface Version: %v\n", d.DescriptorInterface)
	fmt.Printf("    Capability:\n")
	fmt.Printf("      EEPROM: %v\n", d.DescriptorCapability.EEPROM())
	fmt.Printf("      FPGA Configuration: %v\n", d.DescriptorCapability.FPGAConfiguration())
	fmt.Printf("      Flash Memory: %v\n", d.DescriptorCapability.FlashMemory())
	fmt.Printf("      Debug Helper: %v\n", d.DescriptorCapability.DebugHelper())
	fmt.Printf("      XMEGA: %v\n", d.DescriptorCapability.XMEGA())
	fmt.Printf("      High Speed FPGA Configuration: %v\n", d.DescriptorCapability.HighSpeedFPGAConfiguration())
	fmt.Printf("      MAC EEPROM: %v\n", d.DescriptorCapability.MACEEPROM())
	fmt.Printf("      MultiFPGA: %v\n", d.DescriptorCapability.MultiFPGA())
	fmt.Printf("      Temperature Sensor: %v\n", d.DescriptorCapability.TemperatureSensor())
	fmt.Printf("      Flash Memory 2: %v\n", d.DescriptorCapability.FlashMemory2())
	fmt.Printf("      FX3 Firmware: %v\n", d.DescriptorCapability.FX3Firmware())
	fmt.Printf("      Debug Helper 2: %v\n", d.DescriptorCapability.DebugHelper2())
	fmt.Printf("      Default Firmware: %v\n", d.DescriptorCapability.DefaultFirmware())
	fmt.Printf("    Module: %v\n", d.DescriptorModule)
	fmt.Printf("    Serial Number: %v\n", d.DescriptorSerial)
	fmt.Printf("  Board:\n")
	fmt.Printf("    Type: %v\n", d.BoardType)
	fmt.Printf("    Version: %v\n", d.BoardVersion)
	fmt.Printf("  FPGA:\n")
	fmt.Printf("    Type: %v\n", d.FPGAType)
	fmt.Printf("    Package: %v\n", d.FPGAPackage)
	fmt.Printf("    Grade: %v\n", d.FPGAGrade)
	fmt.Printf("  RAM:\n")
	fmt.Printf("    Type: %v\n", d.RAMType)
	fmt.Printf("    Size: %v\n", d.RAMSize)
	fmt.Printf("  Bitstream:\n")
	fmt.Printf("    Start: %v\n", d.BitstreamStart)
	fmt.Printf("    Capacity: %v\n", d.BitstreamCapacity)
	fmt.Printf("    Size: %v\n", d.BitstreamSize)

	return nil
}

func printFPGA(d *ztex.Device) error {
	fst, err := d.FPGAStatus()
	if err != nil {
		return fmt.Errorf("(*ztex.Device).FPGAStatus: %v", err)
	}

	fmt.Printf("FPGA Status:\n")
	fmt.Printf("  Configured: %v\n", fst.FPGAConfigured)
	fmt.Printf("  Checksum: %v\n", fst.FPGAChecksum)
	fmt.Printf("  Transferred: %v\n", fst.FPGATransferred)
	fmt.Printf("  Init: %v\n", fst.FPGAInit)
	fmt.Printf("  Result: %v\n", fst.FPGAResult)
	fmt.Printf("  Swapped: %v\n", fst.FPGASwapped)

	return nil
}

func printFlash(d *ztex.Device) error {
	fst, err := d.FlashStatus()
	if err != nil {
		return fmt.Errorf("(*ztex.Device).FlashStatus: %v", err)
	}

	fmt.Printf("Flash Status:\n")
	fmt.Printf("  Enabled: %v\n", fst.FlashEnabled)
	fmt.Printf("  Sector: %v\n", fst.FlashSector)
	fmt.Printf("  Count: %v\n", fst.FlashCount)
	fmt.Printf("  Error: %v\n", fst.FlashError)

	return nil
}

func main() {
	getopt.Parse()
	if *helpFlag {
		getopt.Usage()
		return
	}

	ctx := gousb.NewContext()
	defer ctx.Close()

	d, err := ztex.OpenDevice(ctx)
	if err != nil {
		log.Fatalf("ztex.OpenDevice: %v", err)
	}
	defer d.Close()

	if *allFlag || *usbFlag {
		if err := printUSB(d.Device); err != nil {
			log.Fatalf("printUSB: %v", err)
		}
	}

	if *allFlag || *ztexFlag {
		if err := printZTEX(d); err != nil {
			log.Fatalf("printZTEX: %v", err)
		}
	}

	if *allFlag || *fpgaFlag {
		if err := printFPGA(d); err != nil {
			log.Fatalf("printFPGA: %v", err)
		}
	}

	if *allFlag || *flashFlag {
		if err := printFlash(d); err != nil {
			log.Fatalf("printFlash: %v", err)
		}
	}
}
