package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_uuid_generate_random(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	v3 = m.Env.Pgmem_random_bytes(m, l0, int32(16))
	mBase = m.M
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	v8 = v4&int32(15) | int32(64)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)) = uint8(v8)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	v14 = v10&int32(63) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v14)
	return
}
func F_uuid_generate_v1mc(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v107 int64
	_ = v107
	var v108 int32
	_ = v108
	v4 = m.G0
	v6 = v4 - int32(128)
	m.G0 = v6
	F_uuid_generate_random(m, v6)
	mBase = m.M
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+10)))
	v11 = v9 | int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+10)) = uint8(v11)
	F_uuid_unparse(m, v6, v6+int32(16))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v20 = v6 - int32(-64)
	F_uuid_generate_time(m, v20)
	mBase = m.M
	v23 = v6 + int32(80)
	F_uuid_unparse(m, v20, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = v6 + int32(103)
	v29 = v6 + int32(40)
	if (v29^v27)&int32(3) != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v107 = F_DirectFunctionCall1Coll(m, int32(3591), int32(0), base.I64_extend_i32_u(v23))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L25
	}
L5:
	;
	goto L4
L6:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v84))) = uint8(v83)
	if v83&int32(255) == int32(0) {
		goto L5
	} else {
		goto L21
	}
L7:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	v82 = v29
	v83 = v35
	v84 = v27
	goto L6
L8:
	;
	goto L9
L9:
	;
	if v29&int32(3) != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v39 = v29
	v41 = v27
	goto L13
L11:
	;
	v53 = v29
	v55 = v27
	goto L12
L12:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v60 = int32(-2139062144)
	if (int32(16843008)-v57|v57)&v60 != v60 {
		v82 = v53
		v83 = v57
		v84 = v55
		goto L6
	} else {
		goto L17
	}
L13:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v42)
	if v42 == int32(0) {
		goto L5
	} else {
		goto L15
	}
L14:
	;
	v53 = v49
	v55 = v47
	goto L12
L15:
	;
	v46 = int32(1)
	v47 = v41 + v46
	v49 = v39 + v46
	if v49&int32(3) != 0 {
		v39 = v49
		v41 = v47
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v65 = v53
	v66 = v57
	v67 = v55
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v66
	v69 = int32(4)
	v70 = v67 + v69
	v72 = v65 + v69
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v77 = int32(-2139062144)
	if (int32(16843008)-v74|v74)&v77 == v77 {
		v65 = v72
		v66 = v74
		v67 = v70
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v82 = v72
	v83 = v74
	v84 = v70
	goto L6
L20:
	;
	goto L19
L21:
	;
	v91 = v82
	v93 = v84
	goto L22
L22:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v93)+1)) = uint8(v94)
	v96 = int32(1)
	if v94 != 0 {
		v91 = v91 + v96
		v93 = v93 + v96
		goto L22
	} else {
		goto L24
	}
L23:
	;
	goto L5
L24:
	;
	goto L23
L25:
	;
	m.G0 = v6 + int32(128)
	return v107
}
func F_uuid_generate_v3(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14400(m, l0, int32(3))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_uuid_generate_v5(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14400(m, l0, int32(5))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_uuid_gt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v9 int64
	_ = v9
	var v11 int64
	_ = v11
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v78 int64
	_ = v78
	var v81 int64
	_ = v81
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	var v89 int64
	_ = v89
	var v91 int64
	_ = v91
	var v93 int64
	_ = v93
	var v95 int64
	_ = v95
	var v116 int64
	_ = v116
	var v117 int64
	_ = v117
	var v152 int64
	_ = v152
	var v154 int64
	_ = v154
	var v155 int64
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v7 = int64(56)
	v9 = int64(65280)
	v11 = int64(40)
	v14 = int64(16711680)
	v16 = int64(24)
	v18 = int64(4278190080)
	v20 = int64(8)
	v41 = v6<<(uint(v7)%64) | v6&v9<<(uint(v11)%64) | (v6&v14<<(uint(v16)%64) | v6&v18<<(uint(v20)%64)) | (int64(base.Ui64(v6)>>(uint(v20)%64))&v18 | int64(base.Ui64(v6)>>(uint(v16)%64))&v14 | (int64(base.Ui64(v6)>>(uint(v11)%64))&v9 | int64(base.Ui64(v6)>>(uint(v7)%64))))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
	v78 = v43<<(uint(v7)%64) | v43&v9<<(uint(v11)%64) | (v43&v14<<(uint(v16)%64) | v43&v18<<(uint(v20)%64)) | (int64(base.Ui64(v43)>>(uint(v20)%64))&v18 | int64(base.Ui64(v43)>>(uint(v16)%64))&v14 | (int64(base.Ui64(v43)>>(uint(v11)%64))&v9 | int64(base.Ui64(v43)>>(uint(v7)%64))))
	if v41 == v78 {
		v81 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
		v82 = int64(56)
		v84 = int64(65280)
		v86 = int64(40)
		v89 = int64(16711680)
		v91 = int64(24)
		v93 = int64(4278190080)
		v95 = int64(8)
		v116 = v81<<(uint(v82)%64) | v81&v84<<(uint(v86)%64) | (v81&v89<<(uint(v91)%64) | v81&v93<<(uint(v95)%64)) | (int64(base.Ui64(v81)>>(uint(v95)%64))&v93 | int64(base.Ui64(v81)>>(uint(v91)%64))&v89 | (int64(base.Ui64(v81)>>(uint(v86)%64))&v84 | int64(base.Ui64(v81)>>(uint(v82)%64))))
		v117 = *(*int64)(unsafe.Add(mBase, uint32(v42)+8))
		v152 = v117<<(uint(v82)%64) | v117&v84<<(uint(v86)%64) | (v117&v89<<(uint(v91)%64) | v117&v93<<(uint(v95)%64)) | (int64(base.Ui64(v117)>>(uint(v95)%64))&v93 | int64(base.Ui64(v117)>>(uint(v91)%64))&v89 | (int64(base.Ui64(v117)>>(uint(v86)%64))&v84 | int64(base.Ui64(v117)>>(uint(v82)%64))))
		if v116 == v152 {
			v162 = int32(0)
		} else {
			v154 = v152
			v155 = v116
			if base.Ui64(v155) < base.Ui64(v154) {
				v159 = int32(-1)
			} else {
				v159 = int32(1)
			}
			v162 = v159
		}
	} else {
		v154 = v78
		v155 = v41
		if base.Ui64(v155) < base.Ui64(v154) {
			v159 = int32(-1)
		} else {
			v159 = int32(1)
		}
		v162 = v159
	}
	return base.I64_extend_i32_u(base.B2i32(int32(0) < v162))
}
func F_uuid_ns_dns(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14401(m, l0, int32(_a_F_uuid_ns_dns_0), int32(_a_F_uuid_ns_dns_1), int32(_a_F_uuid_ns_dns_2), int32(_a_F_uuid_ns_dns_3), int32(_a_F_uuid_ns_dns_4))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_uuid_ns_url(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14401(m, l0, int32(_a_F_uuid_ns_url_0), int32(_a_F_uuid_ns_url_1), int32(_a_F_uuid_ns_url_2), int32(_a_F_uuid_ns_url_3), int32(_a_F_uuid_ns_url_4))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
