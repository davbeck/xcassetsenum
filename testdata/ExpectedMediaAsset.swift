import UIKit

// this class is used to find the bundle the file is in
private class MediaAssetClass: NSObject {}

extension UIImage {
	internal enum MediaAsset: String {
		case Comment = "comment"
		case HomeButton = "HomeButton"
		case More = "more"
		case TrashActivity = "trash-activity"
		case WelcomePlaceholder = "welcome_placeholder"
	}
    
	internal convenience init!(mediaAsset: MediaAsset compatibleWithTraitCollection: UITraitCollection? = nil) {
		self.init(named: mediaAsset.rawValue, inBundle: NSBundle(forClass: MediaAssetClass.self), compatibleWithTraitCollection: compatibleWithTraitCollection)
	}
}
