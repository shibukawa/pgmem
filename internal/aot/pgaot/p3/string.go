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
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if v5 == v10 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v5 == v12 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
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
func F_string_agg_deserialize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = int32(1)
		v18 = v13 + v17
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
		v23 = v21 & v17
		if v23 != 0 {
			v24 = v18
		} else {
			v24 = v13 + int32(4)
		}
		if v21 == int32(1) {
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
			if v30 == int32(18) {
				v33 = int32(16)
			} else {
				v33 = int32(0)
			}
			if base.Ui32((v30-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v40 = int32(4)
			} else {
				v40 = v33
			}
			v51 = v40
		} else {
			v41 = int32(1)
			if v23 != 0 {
				v51 = int32(base.Ui32(v21)>>(uint(v41)%32)) - v41
			} else {
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		*(*int64)(unsafe.Add(mBase, uint32(v10)+20)) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v51
		*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v24
		v57 = v10 + int32(28)
		v58 = int32(0)
		v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v59 == v58 {
			v76 = int32(0)
			if v57 == v76 {
				v84 = v76
			} else {
				v79 = v76
				v80 = v58
				*(*int32)(unsafe.Add(mBase, uint32(v57))) = v79
				v84 = v80
			}
			v87 = v84
		} else {
			v62 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
			switch v62 - int32(429) {
			case 0:
				if v57 == int32(0) {
					v87 = int32(1)
				} else {
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v59)+168))
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
					v79 = v69
					v80 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v57))) = v79
					v84 = v80
					v87 = v84
				}
			case 1:
				if v57 == int32(0) {
					v87 = int32(2)
				} else {
					v74 = *(*int32)(unsafe.Add(mBase, uint32(v59)+368))
					v79 = v74
					v80 = int32(2)
					*(*int32)(unsafe.Add(mBase, uint32(v57))) = v79
					v84 = v80
					v87 = v84
				}
			default:
				v76 = int32(0)
				if v57 == v76 {
					v84 = v76
				} else {
					v79 = v76
					v80 = v58
					*(*int32)(unsafe.Add(mBase, uint32(v57))) = v79
					v84 = v80
				}
				v87 = v84
			}
		}
		if v87 != 0 {
			v88 = int32(_a_F_string_agg_deserialize_0)
			v89 = *(*int32)(unsafe.Add(mBase, _c_F_string_agg_deserialize[0]))
			v91 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
			*(*int32)(unsafe.Add(mBase, _c_F_string_agg_deserialize[0])) = v91
			v93 = F_makeStringInfo(m)
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_string_agg_deserialize[0])) = v89
				v98 = v10 + int32(12)
				v100 = F_pq_getmsgint(m, v98, int32(4))
				mBase = m.M
				v101 = m.ExcPending
				if v101 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v93)+12)) = v100
					v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
					if v103 == int32(1) {
						v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
						if v109 == int32(18) {
							v112 = int32(16)
						} else {
							v112 = int32(0)
						}
						if base.Ui32((v109-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v119 = int32(4)
						} else {
							v119 = v112
						}
						v132 = v119
					} else {
						v120 = int32(1)
						if v103&v120 != 0 {
							v132 = int32(base.Ui32(v103)>>(uint(v120)%32)) - v120
						} else {
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
							v132 = int32(base.Ui32(v126)>>(uint(int32(2))%32)) - int32(4)
						}
					}
					v134 = v132 - int32(4)
					v135 = F_pq_getmsgbytes(m, v98, v134)
					mBase = m.M
					v136 = m.ExcPending
					if v136 != 0 {
						return int32(0)
					} else {
						F_appendBinaryStringInfo(m, v93, v135, v134)
						mBase = m.M
						v138 = m.ExcPending
						if v138 != 0 {
							return int32(0)
						} else {
							F_pq_getmsgend(m, v10+int32(12))
							mBase = m.M
							v142 = m.ExcPending
							if v142 != 0 {
								return int32(0)
							} else {
								m.G0 = v10 + int32(32)
								return v93
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v150 = m.ExcPending
			if v150 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_string_agg_deserialize_1), int32(0))
				mBase = m.M
				v154 = m.ExcPending
				if v154 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_string_agg_deserialize_2), int32(_a_F_string_agg_deserialize_3), int32(_a_F_string_agg_deserialize_4))
					mBase = m.M
					v159 = m.ExcPending
					if v159 != 0 {
						return int32(0)
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
