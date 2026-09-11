package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tsquery_le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_copy(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = l0 + int32(28)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v14 = F_pg_detoast_datum_copy(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = F_CompareTSQ(m, v7, v14)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v18 != v7 {
					F_pfree(m, v7)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
						if v22 != v14 {
							F_pfree(m, v14)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return int32(0)
							} else {
								return base.B2i32(v16 <= int32(0))
							}
						} else {
							return base.B2i32(v16 <= int32(0))
						}
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					if v22 != v14 {
						F_pfree(m, v14)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v16 <= int32(0))
						}
					} else {
						return base.B2i32(v16 <= int32(0))
					}
				}
			}
		}
	}
}
func F_tsquery_ne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_copy(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = l0 + int32(28)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v14 = F_pg_detoast_datum_copy(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = F_CompareTSQ(m, v7, v14)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v18 != v7 {
					F_pfree(m, v7)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
						if v22 != v14 {
							F_pfree(m, v14)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return int32(0)
							} else {
								return base.B2i32(v16 != int32(0))
							}
						} else {
							return base.B2i32(v16 != int32(0))
						}
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					if v22 != v14 {
						F_pfree(m, v14)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v16 != int32(0))
						}
					} else {
						return base.B2i32(v16 != int32(0))
					}
				}
			}
		}
	}
}
func F_tsquery_not(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
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
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_copy(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
		if v10 == int32(0) {
			return v6
		} else {
			v15 = F_palloc0(m, int32(24))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v17 | int32(1)
				v22 = F_palloc0(m, int32(12))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v15))) = v22
					v25 = int32(2)
					*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v25)
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
					v28 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)) = uint8(v28)
					v31 = F_palloc0(m, int32(4))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v31
						v35 = v6 + int32(8)
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
						v40 = F_QT2QTN(m, v35, v35+v36*int32(12))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v42))) = v40
							*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = int32(1)
							v46 = F_QTN2QT(m, v15)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_QTNFree(m, v15)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									if v50 != v6 {
										F_pfree(m, v6)
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
											return int32(0)
										} else {
											return v46
										}
									} else {
										return v46
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
func F_tsquery_phrase(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = F_DirectFunctionCall3Coll(m, int32(1527), int32(0), v4, v5, int32(1))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
