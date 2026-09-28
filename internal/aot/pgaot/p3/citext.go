package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_citext_gt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v14 = F_citextcmp(m, v6, v11, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int64(0)
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v16 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return int64(0)
					} else {
						v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v20 != v11 {
							F_pfree(m, v11)
							mBase = m.M
							v23 = m.ExcPending
							if v23 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_u(base.B2i32(int32(0) < v14))
							}
						} else {
							return base.I64_extend_i32_u(base.B2i32(int32(0) < v14))
						}
					}
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v20 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(base.B2i32(int32(0) < v14))
						}
					} else {
						return base.I64_extend_i32_u(base.B2i32(int32(0) < v14))
					}
				}
			}
		}
	}
}
func F_citext_larger(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v10 = F_pg_detoast_datum_packed(m, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v13 = F_citextcmp(m, v5, v10, v12)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int64(0)
			} else {
				if int32(0) < v13 {
					v17 = v5
				} else {
					v17 = v10
				}
				return base.I64_extend_i32_u(v17)
			}
		}
	}
}
func F_citext_le(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v14 = F_citextcmp(m, v6, v11, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int64(0)
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v16 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return int64(0)
					} else {
						v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v20 != v11 {
							F_pfree(m, v11)
							mBase = m.M
							v23 = m.ExcPending
							if v23 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_u(base.B2i32(v14 <= int32(0)))
							}
						} else {
							return base.I64_extend_i32_u(base.B2i32(v14 <= int32(0)))
						}
					}
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v20 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(base.B2i32(v14 <= int32(0)))
						}
					} else {
						return base.I64_extend_i32_u(base.B2i32(v14 <= int32(0)))
					}
				}
			}
		}
	}
}
func F_citext_ne(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v15 = int32(1)
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	v19 = v17 & v15
	if v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v20 = v15
	goto L6
L5:
	;
	v20 = int32(4)
	goto L6
L6:
	;
	if v17 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v50 = F_str_tolower(m, v8+v20, v48, int32(100))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L18
	}
L8:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
	if v27 == int32(18) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v38 = int32(1)
	if v19 != 0 {
		v48 = int32(base.Ui32(v17)>>(uint(v38)%32)) - v38
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v30 = int32(16)
	goto L13
L12:
	;
	v30 = int32(0)
	goto L13
L13:
	;
	if base.Ui32((v27-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v37 = int32(4)
	goto L16
L15:
	;
	v37 = v30
	goto L16
L16:
	;
	v48 = v37
	goto L7
L17:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v48 = int32(base.Ui32(v42)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	v52 = int32(1)
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	v56 = v54 & v52
	if v56 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v57 = v52
	goto L21
L20:
	;
	v57 = int32(4)
	goto L21
L21:
	;
	if v54 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v87 = F_str_tolower(m, v13+v57, v85, int32(100))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L33
	}
L23:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	if v64 == int32(18) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v75 = int32(1)
	if v56 != 0 {
		v85 = int32(base.Ui32(v54)>>(uint(v75)%32)) - v75
		goto L22
	} else {
		goto L32
	}
L26:
	;
	v67 = int32(16)
	goto L28
L27:
	;
	v67 = int32(0)
	goto L28
L28:
	;
	if base.Ui32((v64-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v74 = int32(4)
	goto L31
L30:
	;
	v74 = v67
	goto L31
L31:
	;
	v85 = v74
	goto L22
L32:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v85 = int32(base.Ui32(v79)>>(uint(int32(2))%32)) - int32(4)
	goto L22
L33:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if base.B2i32(v91 == int32(0))|base.B2i32(v91 != v94) != 0 {
		v112 = v91
		v113 = v94
		goto L35
	} else {
		goto L36
	}
L34:
	;
	F_pfree(m, v50)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L41
	}
L35:
	;
	goto L34
L36:
	;
	v97 = v50
	v98 = v87
	goto L37
L37:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+1)))
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
	if v102 == int32(0) {
		v112 = v102
		v113 = v101
		goto L35
	} else {
		goto L39
	}
L38:
	;
	v112 = v102
	v113 = v101
	goto L35
L39:
	;
	v105 = int32(1)
	if v102 == v101 {
		v97 = v97 + v105
		v98 = v98 + v105
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	F_pfree(m, v87)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v119 != v8 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	F_pfree(m, v8)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v123 != v13 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L45
L47:
	;
	F_pfree(m, v13)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	return base.I64_extend_i32_u(base.B2i32(v112-v113 != int32(0)))
L50:
	;
	goto L49
}
func F_citext_pattern_cmp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			v13 = F_internal_citext_pattern_cmp(m, v6, v11)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int64(0)
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v15 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return int64(0)
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v19 != v11 {
							F_pfree(m, v11)
							mBase = m.M
							v22 = m.ExcPending
							if v22 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_s(v13)
							}
						} else {
							return base.I64_extend_i32_s(v13)
						}
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v19 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_s(v13)
						}
					} else {
						return base.I64_extend_i32_s(v13)
					}
				}
			}
		}
	}
}
func F_citext_pattern_lt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			v13 = F_internal_citext_pattern_cmp(m, v6, v11)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int64(0)
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v15 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return int64(0)
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v19 != v11 {
							F_pfree(m, v11)
							mBase = m.M
							v22 = m.ExcPending
							if v22 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_u(int32(base.Ui32(v13) >> (uint(int32(31)) % 32)))
							}
						} else {
							return base.I64_extend_i32_u(int32(base.Ui32(v13) >> (uint(int32(31)) % 32)))
						}
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v19 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(int32(base.Ui32(v13) >> (uint(int32(31)) % 32)))
						}
					} else {
						return base.I64_extend_i32_u(int32(base.Ui32(v13) >> (uint(int32(31)) % 32)))
					}
				}
			}
		}
	}
}
