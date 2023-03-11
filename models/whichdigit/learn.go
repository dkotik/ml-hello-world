package whichdigit

func (d *Detector) Learn(image []byte, label Digit) (err error) {
	if err = d.IsDetectableImage(image); err != nil {
		return err
	}
	d.activateUsingImage(image)
	return nil
}
