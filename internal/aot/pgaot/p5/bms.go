package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bms_equal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	v3 = int32(0)
	if base.B2i32(l0 == v3)|base.B2i32(l1 == v3) != 0 {
		v49 = base.B2i32(l0|l1 == v3)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v49
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v17 != v18 {
		v49 = int32(0)
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = int32(1)
	if v17 <= v20 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v23 = v20
	goto L6
L5:
	;
	v23 = v17
	goto L6
L6:
	;
	v24 = int32(8)
	v29 = int32(0)
	goto L7
L7:
	;
	v37 = v29 << (uint(int32(2)) % 32)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0+v24+v37)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1+v24+v37)))
	v42 = base.B2i32(v39 == v41)
	if v39 != v41 {
		v49 = v42
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v49 = v42
	goto L1
L9:
	;
	v45 = v29 + int32(1)
	if v45 != v23 {
		v29 = v45
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
}
func F_bms_membership(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	v2 = int32(0)
	if l0 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v11 = int32(1)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v12 <= v11 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v15 = v11
	goto L6
L5:
	;
	v15 = v12
	goto L6
L6:
	;
	v19 = int32(0)
	v21 = v2
	goto L7
L7:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(8)+v19<<(uint(int32(2))%32))))
	if v28 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	return v40
L9:
	;
	goto L8
L10:
	;
	v29 = int32(2)
	if v21 != 0 {
		v40 = v29
		goto L9
	} else {
		goto L13
	}
L11:
	;
	v35 = v21
	goto L12
L12:
	;
	v37 = v19 + int32(1)
	if v37 != v15 {
		v19 = v37
		v21 = v35
		goto L7
	} else {
		goto L15
	}
L13:
	;
	v30 = int32(1)
	if base.Ui32(v30) < base.Ui32(base.I32_popcnt(v28)) {
		v40 = v29
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v35 = v30
	goto L12
L15:
	;
	v40 = v35
	goto L9
}
func F_bms_num_members(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v36 int64
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int64
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v66 int32
	_ = v66
	var v67 int64
	_ = v67
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v73 int64
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int64
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int64
	_ = v88
	var v92 int32
	_ = v92
	var v93 int64
	_ = v93
	var v94 int64
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int64
	_ = v102
	var v112 int64
	_ = v112
	var v115 int64
	_ = v115
	v3 = int64(0)
	if l0 == int32(0) {
		return int32(0)
	} else {
		v9 = l0 + int32(8)
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v10 == int32(1) {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
			return base.I32_popcnt(v13)
		} else {
			v17 = v10 << (uint(int32(2)) % 32)
			if v17 <= int32(7) {
				if v17 == int32(0) {
					v115 = v3
				} else {
					v22 = v17
					v23 = v9
					v24 = v3
					for {
						v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+3)))
						v26 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_bms_num_members[0]))))
						v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+2)))
						v28 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_bms_num_members[0]))))
						v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
						v30 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v29)+uint32(_c_F_bms_num_members[0]))))
						v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
						v32 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v31)+uint32(_c_F_bms_num_members[0]))))
						v36 = v26 + (v28 + (v30 + (v24 + v32)))
						v37 = int32(4)
						v40 = v22 - v37
						if v40 != 0 {
							v22 = v40
							v23 = v23 + v37
							v24 = v36
							continue
						} else {
							break
						}
						break
					}
					v115 = v36
				}
			} else {
				v41 = int64(0)
				v42 = int32(0)
				if v17 == v42 {
					v112 = int64(0)
				} else {
					v49 = v17 & int32(3)
					if base.Ui32(int32(4)) <= base.Ui32(v17) {
						v54 = v9
						v56 = v41
						v59 = v42
						for {
							v60 = int32(4)
							v61 = v54 + v60
							v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+3)))
							v63 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v62)+uint32(_c_F_bms_num_members[0]))))
							v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+2)))
							v65 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v64)+uint32(_c_F_bms_num_members[0]))))
							v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+1)))
							v67 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v66)+uint32(_c_F_bms_num_members[0]))))
							v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
							v69 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v68)+uint32(_c_F_bms_num_members[0]))))
							v73 = v63 + (v65 + (v67 + (v56 + v69)))
							v75 = v59 + v60
							if v75 != v17&int32(-4) {
								v54 = v61
								v56 = v73
								v59 = v75
								continue
							} else {
								break
							}
							break
						}
						if v49 == int32(0) {
							v102 = v73
						} else {
							v79 = v61
							v81 = v73
							v86 = v79
							v87 = int32(0)
							v88 = v81
							for {
								v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
								v93 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v92)+uint32(_c_F_bms_num_members[0]))))
								v94 = v88 + v93
								v95 = int32(1)
								v98 = v87 + v95
								if v98 != v49 {
									v86 = v86 + v95
									v87 = v98
									v88 = v94
									continue
								} else {
									break
								}
								break
							}
							v102 = v94
						}
					} else {
						v79 = v9
						v81 = v41
						v86 = v79
						v87 = int32(0)
						v88 = v81
						for {
							v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
							v93 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v92)+uint32(_c_F_bms_num_members[0]))))
							v94 = v88 + v93
							v95 = int32(1)
							v98 = v87 + v95
							if v98 != v49 {
								v86 = v86 + v95
								v87 = v98
								v88 = v94
								continue
							} else {
								break
							}
							break
						}
						v102 = v94
					}
					v112 = v102
				}
				v115 = v112
			}
			return base.I32_wrap_i64(v115)
		}
	}
}
func F_bms_subset_compare(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return base.B2i32(l1 != int32(0))
L2:
	;
	goto L3
L3:
	;
	if l1 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(2)
L5:
	;
	goto L6
L6:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v21 < v22 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v24 = v21
	goto L9
L8:
	;
	v24 = v22
	goto L9
L9:
	;
	if v24 <= int32(1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v27 = int32(1)
	goto L12
L11:
	;
	v27 = v24
	goto L12
L12:
	;
	v28 = int32(8)
	v32 = int32(0)
	v34 = v32
	v35 = v32
	goto L15
L13:
	;
	return int32(3)
L14:
	;
	return v86
L15:
	;
	v45 = v35 << (uint(int32(2)) % 32)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0+v28+v45)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45+(l1+v28))))
	if v47&(v49^int32(-1)) != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	if v22 < v21 {
		goto L25
	} else {
		goto L26
	}
L17:
	;
	v71 = v35 + int32(1)
	if v71 != v27 {
		v34 = v69
		v35 = v71
		goto L15
	} else {
		goto L24
	}
L18:
	;
	if base.B2i32(v34 == int32(1))|v49&(v47^int32(-1)) != 0 {
		v86 = int32(3)
		goto L14
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v49&(v47^int32(-1)) == int32(0) {
		v69 = v34
		goto L17
	} else {
		goto L22
	}
L21:
	;
	v69 = int32(2)
	goto L17
L22:
	;
	if v34 == int32(2) {
		goto L13
	} else {
		goto L23
	}
L23:
	;
	v69 = int32(1)
	goto L17
L24:
	;
	goto L16
L25:
	;
	if v69 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	if v22 <= v21 {
		v86 = v69
		goto L14
	} else {
		goto L31
	}
L28:
	;
	v78 = int32(3)
	goto L30
L29:
	;
	v78 = int32(2)
	goto L30
L30:
	;
	return v78
L31:
	;
	if v69 == int32(2) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v85 = int32(3)
	goto L34
L33:
	;
	v85 = int32(1)
	goto L34
L34:
	;
	v86 = v85
	goto L14
}
