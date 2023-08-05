# Astroalign

## Use cases

* I have two images and I want one of them transformed to look like the other
* I have two images and I want to know the transformation between the two
	* I need the matrix
	* I need the parameters
* I have an image and I want to identify each star on the other image
* I already have star correspondences and want to know what is the transformation
* I have a set of points and I want to know where they lay on my other image
* I want to use a different invariant map from the one provided


### Case 1: I have two images and I want one of them transformed to look like a reference

Option 1 (current)

    >>> import astroalign as aa
    >>> img_transf = aa.align_image(ref_image, img2transf, **kwargs)

Option 2

    >>> img_transf = aa.register(reference=ref_image, target=img2transform, **kwargs)


### Case 2: I have two images and I want to know the transformation between the two

Option 1

    >>> import astroalign as aa
    >>> transf_matrix, params = aa.find_transform(transform=img2transf, reference=ref_image, **kwargs)

Option 2

    >>> import astroalign as aa
    >>> transf_matrix, params = aa.mapping(target, source, **kwargs)

Option 3

    >>> import astroalign as aa
    >>> transf_matrix, params, control_matches = aa.mapping(target, source, **kwargs)

  
### Case 3: I have an image and I want to identify each star on the other image

Same as before, transformation and control point matches are done simultaneously

### Case 4: I already have star correspondences and want to know what is the transformation

Use scikit-image. Maybe a wrapper.

### Case 5: I have a set of points and I want to know where they lay on my other image

Just mimic scikit-image here?

### Case 6: I want to use a different invariant map from the one provided

