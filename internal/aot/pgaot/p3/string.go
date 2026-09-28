package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_appendStringInfo(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_appendStringInfo[0]))
	goto L1
L1:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_appendStringInfo[0])) = v14
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l2
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v29 = v27 - v28
	if int32(16) <= v29 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v46 + v35
	m.G0 = v11 + int32(16)
	return
L3:
	;
	goto L2
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v35 = F_pvsnprintf(m, v32+v28, v29, l1, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v43 = int32(32)
	goto L6
L6:
	;
	F_enlargeStringInfo(m, l0, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L7
	} else {
		goto L10
	}
L7:
	;
	return
L8:
	;
	if base.Ui32(v35) < base.Ui32(v29) {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v41 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v38+v39))) = uint8(v41)
	v43 = v35
	goto L6
L10:
	;
	goto L1
}
func F_makeStringInfo(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	v4 = F_palloc(m, int32(16))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v9 = F_palloc(m, int32(1024))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = int32(1024)
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = v9
			v14 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v14)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = v14
			*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = v14
			return v4
		}
	}
}
func F_set_string_field(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = l2
	if v5 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if v5 == v10 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v5 == v12 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v5 == v14 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v19 = l0 + int32(56)
	goto L6
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v22 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	F_pfree(m, v5)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	if v5 == v23 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	goto L7
L11:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	if v5 != v25 {
		v19 = v22
		goto L6
	} else {
		goto L12
	}
L12:
	;
	goto L1
L13:
	;
	return
L14:
	;
	goto L1
}
func F_string_agg_deserialize(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = int32(1)
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
		v20 = v18 & v16
		if v20 != 0 {
			v21 = v16
		} else {
			v21 = int32(4)
		}
		if v18 == int32(1) {
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
			if v28 == int32(18) {
				v31 = int32(16)
			} else {
				v31 = int32(0)
			}
			if base.Ui32((v28-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v38 = int32(4)
			} else {
				v38 = v31
			}
			v49 = v38
		} else {
			v39 = int32(1)
			if v20 != 0 {
				v49 = int32(base.Ui32(v18)>>(uint(v39)%32)) - v39
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v49 = int32(base.Ui32(v43)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		*(*int64)(unsafe.Add(mBase, uint32(v9)+20)) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v49
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v12 + v21
		v55 = v9 + int32(28)
		v56 = int32(0)
		v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v57 == v56 {
			v74 = int32(0)
			if v55 == v74 {
				v82 = v74
			} else {
				v77 = v74
				v78 = v56
				*(*int32)(unsafe.Add(mBase, uint32(v55))) = v77
				v82 = v78
			}
			v85 = v82
		} else {
			v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
			switch v60 - int32(435) {
			case 0:
				if v55 == int32(0) {
					v85 = int32(1)
				} else {
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v57)+168))
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+20))
					v77 = v67
					v78 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v55))) = v77
					v82 = v78
					v85 = v82
				}
			case 1:
				if v55 == int32(0) {
					v85 = int32(2)
				} else {
					v72 = *(*int32)(unsafe.Add(mBase, uint32(v57)+376))
					v77 = v72
					v78 = int32(2)
					*(*int32)(unsafe.Add(mBase, uint32(v55))) = v77
					v82 = v78
					v85 = v82
				}
			default:
				v74 = int32(0)
				if v55 == v74 {
					v82 = v74
				} else {
					v77 = v74
					v78 = v56
					*(*int32)(unsafe.Add(mBase, uint32(v55))) = v77
					v82 = v78
				}
				v85 = v82
			}
		}
		if v85 != 0 {
			v86 = int32(_a_F_string_agg_deserialize_0)
			v87 = *(*int32)(unsafe.Add(mBase, _c_F_string_agg_deserialize[0]))
			v89 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
			*(*int32)(unsafe.Add(mBase, _c_F_string_agg_deserialize[0])) = v89
			v91 = F_makeStringInfo(m)
			mBase = m.M
			v92 = m.ExcPending
			if v92 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_string_agg_deserialize[0])) = v87
				v96 = v9 + int32(12)
				v98 = F_pq_getmsgint(m, v96, int32(4))
				mBase = m.M
				v99 = m.ExcPending
				if v99 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v91)+12)) = v98
					v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
					if v101 == int32(1) {
						v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
						if v107 == int32(18) {
							v110 = int32(16)
						} else {
							v110 = int32(0)
						}
						if base.Ui32((v107-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v117 = int32(4)
						} else {
							v117 = v110
						}
						v130 = v117
					} else {
						v118 = int32(1)
						if v101&v118 != 0 {
							v130 = int32(base.Ui32(v101)>>(uint(v118)%32)) - v118
						} else {
							v124 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
							v130 = int32(base.Ui32(v124)>>(uint(int32(2))%32)) - int32(4)
						}
					}
					v132 = v130 - int32(4)
					v133 = F_pq_getmsgbytes(m, v96, v132)
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return int64(0)
					} else {
						F_appendBinaryStringInfo(m, v91, v133, v132)
						mBase = m.M
						v136 = m.ExcPending
						if v136 != 0 {
							return int64(0)
						} else {
							F_pq_getmsgend(m, v9+int32(12))
							mBase = m.M
							v140 = m.ExcPending
							if v140 != 0 {
								return int64(0)
							} else {
								m.G0 = v9 + int32(32)
								return base.I64_extend_i32_u(v91)
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v149 = m.ExcPending
			if v149 != 0 {
				return int64(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_string_agg_deserialize_1), int32(0))
				mBase = m.M
				v153 = m.ExcPending
				if v153 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_string_agg_deserialize_2), int32(_a_F_string_agg_deserialize_3), int32(_a_F_string_agg_deserialize_4))
					mBase = m.M
					v158 = m.ExcPending
					if v158 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_transform_string_values_array_element_start(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v8 = v6 + v7
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8-int32(1)))))
	if v11 != int32(91) {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		if v14 <= v7+int32(1) {
			F_appendStringInfoChar(m, v5, int32(44))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		} else {
			v25 = int32(44)
			*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v25)
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
			v30 = v28 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v30
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
			v34 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v32+v30))) = uint8(v34)
			return int32(0)
		}
	} else {
		return int32(0)
	}
}
