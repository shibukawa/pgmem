package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_macaddr8_eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v7 = int32(24)
	v9 = int32(65280)
	v11 = int32(8)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v6<<(uint(v7)%32)|v6&v9<<(uint(v11)%32)|(int32(base.Ui32(v6)>>(uint(v11)%32))&v9|int32(base.Ui32(v6)>>(uint(v7)%32))) != v23<<(uint(v7)%32)|v23&v9<<(uint(v11)%32)|(int32(base.Ui32(v23)>>(uint(v11)%32))&v9|int32(base.Ui32(v23)>>(uint(v7)%32))) {
		v76 = v2
	} else {
		v40 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
		v41 = int32(24)
		v43 = int32(65280)
		v45 = int32(8)
		v55 = v40<<(uint(v41)%32) | v40&v43<<(uint(v45)%32) | (int32(base.Ui32(v40)>>(uint(v45)%32))&v43 | int32(base.Ui32(v40)>>(uint(v41)%32)))
		v56 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
		v71 = v56<<(uint(v41)%32) | v56&v43<<(uint(v45)%32) | (int32(base.Ui32(v56)>>(uint(v45)%32))&v43 | int32(base.Ui32(v56)>>(uint(v41)%32)))
		if base.Ui32(v55) < base.Ui32(v71) {
			v76 = v2
		} else {
			v76 = base.B2i32(base.Ui32(v55) <= base.Ui32(v71))
		}
	}
	return v76
}
func F_macaddr8_ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v61 int32
	_ = v61
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v7 = int32(24)
	v9 = int32(65280)
	v11 = int32(8)
	v21 = v6<<(uint(v7)%32) | v6&v9<<(uint(v11)%32) | (int32(base.Ui32(v6)>>(uint(v11)%32))&v9 | int32(base.Ui32(v6)>>(uint(v7)%32)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v38 = v23<<(uint(v7)%32) | v23&v9<<(uint(v11)%32) | (int32(base.Ui32(v23)>>(uint(v11)%32))&v9 | int32(base.Ui32(v23)>>(uint(v7)%32)))
	if base.Ui32(v21) < base.Ui32(v38) {
		return int32(0)
	} else {
		if base.Ui32(v38) < base.Ui32(v21) {
			return int32(1)
		} else {
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
			v46 = int32(24)
			v48 = int32(65280)
			v50 = int32(8)
			v61 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
			return base.B2i32(base.Ui32(v61<<(uint(v46)%32)|v61&v48<<(uint(v50)%32)|(int32(base.Ui32(v61)>>(uint(v50)%32))&v48|int32(base.Ui32(v61)>>(uint(v46)%32)))) <= base.Ui32(v45<<(uint(v46)%32)|v45&v48<<(uint(v50)%32)|(int32(base.Ui32(v45)>>(uint(v50)%32))&v48|int32(base.Ui32(v45)>>(uint(v46)%32)))))
		}
	}
}
func F_macaddr8_lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v61 int32
	_ = v61
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v7 = int32(24)
	v9 = int32(65280)
	v11 = int32(8)
	v21 = v6<<(uint(v7)%32) | v6&v9<<(uint(v11)%32) | (int32(base.Ui32(v6)>>(uint(v11)%32))&v9 | int32(base.Ui32(v6)>>(uint(v7)%32)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v38 = v23<<(uint(v7)%32) | v23&v9<<(uint(v11)%32) | (int32(base.Ui32(v23)>>(uint(v11)%32))&v9 | int32(base.Ui32(v23)>>(uint(v7)%32)))
	if base.Ui32(v21) < base.Ui32(v38) {
		return int32(1)
	} else {
		if base.Ui32(v38) < base.Ui32(v21) {
			return int32(0)
		} else {
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
			v46 = int32(24)
			v48 = int32(65280)
			v50 = int32(8)
			v61 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
			return base.B2i32(base.Ui32(v45<<(uint(v46)%32)|v45&v48<<(uint(v50)%32)|(int32(base.Ui32(v45)>>(uint(v50)%32))&v48|int32(base.Ui32(v45)>>(uint(v46)%32)))) < base.Ui32(v61<<(uint(v46)%32)|v61&v48<<(uint(v50)%32)|(int32(base.Ui32(v61)>>(uint(v50)%32))&v48|int32(base.Ui32(v61)>>(uint(v46)%32)))))
		}
	}
}
func F_macaddr8_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	v11 = m.G0
	v12 = int32(32)
	v13 = v11 - v12
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = F_palloc(m, v12)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
		v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
		v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+2)))
		v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+3)))
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
		v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+5)))
		v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+6)))
		v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+7)))
		*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v28
		*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v27
		*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v26
		*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v25
		*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v13))) = v21
		v39 = F_pg_snprintf(m, v17, int32(32), int32(29543), v13)
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return int32(0)
		} else {
			m.G0 = v13 + int32(32)
			return v17
		}
	}
}
func F_macaddr8_send(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_pq_begintypsend(m, v7)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
		F_enlargeStringInfo(m, v7, int32(1))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			*(*uint8)(unsafe.Add(mBase, uint32(v18+v19))) = uint8(v14)
			v22 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v18 + v22
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
			F_enlargeStringInfo(m, v7, v22)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				*(*uint8)(unsafe.Add(mBase, uint32(v29+v30))) = uint8(v25)
				v33 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v29 + v33
				v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+2)))
				F_enlargeStringInfo(m, v7, v33)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
					*(*uint8)(unsafe.Add(mBase, uint32(v40+v41))) = uint8(v36)
					v44 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v40 + v44
					v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+3)))
					F_enlargeStringInfo(m, v7, v44)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
						*(*uint8)(unsafe.Add(mBase, uint32(v51+v52))) = uint8(v47)
						v55 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v51 + v55
						v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+4)))
						F_enlargeStringInfo(m, v7, v55)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
							*(*uint8)(unsafe.Add(mBase, uint32(v62+v63))) = uint8(v58)
							v66 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v62 + v66
							v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+5)))
							F_enlargeStringInfo(m, v7, v66)
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int32(0)
							} else {
								v73 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
								*(*uint8)(unsafe.Add(mBase, uint32(v73+v74))) = uint8(v69)
								v77 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v73 + v77
								v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+6)))
								F_enlargeStringInfo(m, v7, v77)
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return int32(0)
								} else {
									v84 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
									v85 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
									*(*uint8)(unsafe.Add(mBase, uint32(v84+v85))) = uint8(v80)
									v88 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v84 + v88
									v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+7)))
									F_enlargeStringInfo(m, v7, v88)
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return int32(0)
									} else {
										v95 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
										v96 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
										*(*uint8)(unsafe.Add(mBase, uint32(v95+v96))) = uint8(v91)
										v100 = v95 + int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v100
										v103 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
										*(*int32)(unsafe.Add(mBase, uint32(v103))) = v100 << (uint(int32(2)) % 32)
										m.G0 = v7 + int32(16)
										return v103
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_macaddr8_trunc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_palloc0(m, int32(8))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
		*(*uint8)(unsafe.Add(mBase, uint32(v5))) = uint8(v9)
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+1)))
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+1)) = uint8(v11)
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+2)))
		v14 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+3)) = v14
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+2)) = uint8(v13)
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+7)) = uint8(v14)
		return v5
	}
}
