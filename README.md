# xcassetsenum

Convert an Xcode `.xcassets` catalog file into a Swift enum for type safe compiler checked image assets.

## Usage

Simply pass in the path to your xcassets catalog and optionally the access control you want to use.

```
xcassetsenum --catalog path/to/Media.xcassets --access_control public
```

A swift file will be created or updated in the same directory as the asset catalog with the same name as the catalog. Here's an example of what gets generated:

```swift
import UIKit

// this class is used to find the bundle the file is in
private class MediaAssetClass: NSObject {}

extension UIImage {
	public enum MediaAsset: String {
		case Comment = "comment"
		case HomeButton = "HomeButton"
		case More = "more"
		case TrashActivity = "trash-activity"
		case WelcomePlaceholder = "welcome_placeholder"
	}
    
	public convenience init!(mediaAsset: MediaAsset compatibleWithTraitCollection: UITraitCollection? = nil) {
		self.init(named: mediaAsset.rawValue, inBundle: NSBundle(forClass: MediaAssetClass.self), compatibleWithTraitCollection: compatibleWithTraitCollection)
	}
}
```

You can then create an image using the following:

```swift
let image = UIImage(mediaAsset: .Comment)
```

Note that the result is not optional, unlike `UIImage(name: "comment")`.

If you add xcassetsenum as a run script build phase of your project, anytime you make a change to the catalog the enum values will be updated and you will get a compile time warning if you try to use an asset that has been renamed or removed. This gives you an error if you accidentally remove an asset that is still being used.

### Options

```
  -h, --help                        display help information
  -c, --catalog                    *The path to an xcassets catalog to process into a Swift enum. Required.
  -a, --access_control[=internal]   The access to use for the generated enum, such as public, private, internal. Defaults to internal.
```

## To Do:

- [ ] Document Xcode integration
- [ ] Add `--output` argument to change where the file is saved.
- [ ] Add `--template` argument to change the swift source template.
- [ ] Upload prebuilt binary for macOS.