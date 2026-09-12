package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_johab_to_utf8(m *base.Module, l0 int32) int32 {
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
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(40), int32(6))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		v24 = F_LocalToUtf(m, v6, v10, v5, int32(4339148), v18, v18, v18, int32(40), base.B2i32(v7 != v18))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			return v24
		}
	}
}
func F_jsonpath_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
		F_initStringInfo(m, v6+int32(32))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			F_enlargeStringInfo(m, v6+int32(32), int32(base.Ui32(v13)>>(uint(int32(2))%32)))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
				if int32(0) <= v24 {
					F_appendStringInfoString(m, v6+int32(32), int32(695182))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						F_jspInitByBuffer(m, v6+int32(4), v9+int32(8), int32(0))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							F_printJsonPathItem(m, v6+int32(32), v6+int32(4), int32(0), int32(1))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v6)+32))
								m.G0 = v6 + int32(48)
								return v47
							}
						}
					}
				} else {
					F_jspInitByBuffer(m, v6+int32(4), v9+int32(8), int32(0))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						F_printJsonPathItem(m, v6+int32(32), v6+int32(4), int32(0), int32(1))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v6)+32))
							m.G0 = v6 + int32(48)
							return v47
						}
					}
				}
			}
		}
	}
}
func F_jsonpath_send(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	v3 = m.G0
	v5 = v3 - int32(48)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_initStringInfo(m, v5+int32(4))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			F_enlargeStringInfo(m, v5+int32(4), int32(base.Ui32(v18)>>(uint(int32(2))%32)))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
				if int32(0) <= v23 {
					F_appendStringInfoString(m, v5+int32(4), int32(695182))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						F_jspInitByBuffer(m, v5+int32(20), v8+int32(8), int32(0))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							F_printJsonPathItem(m, v5+int32(4), v5+int32(20), int32(0), int32(1))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								F_pq_begintypsend(m, v5+int32(20))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									F_enlargeStringInfo(m, v5+int32(20), int32(1))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v5)+24))
										v56 = *(*int32)(unsafe.Add(mBase, uint32(v5)+20))
										v58 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v55+v56))) = uint8(v58)
										*(*int32)(unsafe.Add(mBase, uint32(v5)+24)) = v55 + v58
										v65 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
										v66 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
										F_pq_sendtext(m, v5+int32(20), v65, v66)
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return int32(0)
										} else {
											v69 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
											F_pfree(m, v69)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int32(0)
											} else {
												v73 = v5 + int32(20)
												v75 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
												v76 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v75))) = v76 << (uint(int32(2)) % 32)
												m.G0 = v5 + int32(48)
												return v75
											}
										}
									}
								}
							}
						}
					}
				} else {
					F_jspInitByBuffer(m, v5+int32(20), v8+int32(8), int32(0))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						F_printJsonPathItem(m, v5+int32(4), v5+int32(20), int32(0), int32(1))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							F_pq_begintypsend(m, v5+int32(20))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								F_enlargeStringInfo(m, v5+int32(20), int32(1))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v5)+24))
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v5)+20))
									v58 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v55+v56))) = uint8(v58)
									*(*int32)(unsafe.Add(mBase, uint32(v5)+24)) = v55 + v58
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
									v66 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
									F_pq_sendtext(m, v5+int32(20), v65, v66)
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return int32(0)
									} else {
										v69 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
										F_pfree(m, v69)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return int32(0)
										} else {
											v73 = v5 + int32(20)
											v75 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
											v76 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v75))) = v76 << (uint(int32(2)) % 32)
											m.G0 = v5 + int32(48)
											return v75
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
}
